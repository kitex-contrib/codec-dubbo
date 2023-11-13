/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package registry

import (
	"time"

	"github.com/kitex-contrib/codec-dubbo/registries"
)

type Options struct {
	Servers        []string
	RegistryGroup  string
	Username       string
	Password       string
	SessionTimeout time.Duration
}

func (o *Options) Apply(opts []Option) {
	for _, opt := range opts {
		opt.F(o)
	}
}

func newOptions(opts []Option) *Options {
	o := &Options{}

	o.Apply(opts)

	if len(o.Servers) <= 0 {
		panic("Please specify at least one zookeeper server address. e.g. WithServers(\"127.0.0.1:2181\")")
	}

	if o.RegistryGroup == "" {
		o.RegistryGroup = registries.DefaultRegistryGroup
	}
	if o.SessionTimeout == 0 {
		o.SessionTimeout = defaultSessionTimeout
	}
	return o
}

type Option struct {
	F func(o *Options)
}

// WithServers configures target zookeeper servers that zookeeperResolver would connect to.
// Please specify at least one server address, e.g. WithServers("127.0.0.1:2181")
func WithServers(servers ...string) Option {
	return Option{F: func(o *Options) {
		o.Servers = servers
	}}
}

// WithRegistryGroup configures the group of the zookeepers serving the target dubbo Service.
// In dubbo side, this group is referred to RegistryConfig.group.
func WithRegistryGroup(group string) Option {
	return Option{F: func(o *Options) {
		o.RegistryGroup = group
	}}
}

// WithSessionTimeout configures the amount of time for which a session
// is considered valid after losing connection to a server.
// Within the session timeout it's possible to reestablish a connection
// to a different server and keep the same session.
// The default SessionTimeout would be 30 * time.Second
func WithSessionTimeout(timeout time.Duration) Option {
	return Option{F: func(o *Options) {
		o.SessionTimeout = timeout
	}}
}

// WithUsername configures the username to connect to zookeeper servers.
// Please specify this Option with WithPassword together otherwise it
// would not make any sense.
func WithUsername(name string) Option {
	return Option{F: func(o *Options) {
		o.Username = name
	}}
}

// WithPassword configures the password to connect to zookeeper servers.
// Please specify this Option with WithUsername together otherwise it
// would not make any sense.
func WithPassword(password string) Option {
	return Option{F: func(o *Options) {
		o.Password = password
	}}
}
