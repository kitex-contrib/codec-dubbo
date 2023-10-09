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

package dubbo

type Options struct {
	JavaClassName string
}

func (o *Options) Apply(opts []Option) {
	for _, opt := range opts {
		opt.F(o)
	}
}

func newOptions(opts []Option) *Options {
	o := &Options{}

	o.Apply(opts)
	if o.JavaClassName == "" {
		panic("DubboCodec must be initialized with JavaClassName. Please use dubbo.WithJavaClassName().")
	}
	return o
}

type Option struct {
	F func(o *Options)
}

// WithJavaClassName configures InterfaceName for server-side service and specifies target InterfaceName for client.
// Each client and service must have its own corresponding DubboCodec.
func WithJavaClassName(name string) Option {
	return Option{F: func(o *Options) {
		o.JavaClassName = name
	}}
}
