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

package registries

import (
	"net"
	"testing"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/stretchr/testify/assert"
)

func TestURL_FromInfo(t *testing.T) {
	commonAddrFunc := func() net.Addr {
		addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:8888")
		return addr
	}
	commonTags := map[string]string{
		DubboServiceInterfaceKey: "testInterface",
	}
	tests := []struct {
		desc     string
		initAddr func() net.Addr
		info     *registry.Info
		expected func(t *testing.T, u *URL, err error)
	}{
		{
			desc: "Info without port",
			initAddr: func() net.Addr {
				addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0")
				return addr
			},
			info: &registry.Info{
				Tags: commonTags,
			},
			expected: func(t *testing.T, u *URL, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			desc: "Info with port-only ipv4 addr",
			initAddr: func() net.Addr {
				addr, _ := net.ResolveTCPAddr("tcp", ":8888")
				return addr
			},
			info: &registry.Info{
				Tags: commonTags,
			},
			expected: func(t *testing.T, u *URL, err error) {
				assert.Nil(t, err)
				assert.NotEmpty(t, u.host)
			},
		},
		{
			desc: "Info with port-only ipv6 addr",
			initAddr: func() net.Addr {
				addr, _ := net.ResolveTCPAddr("tcp", "[::]:8888")
				return addr
			},
			info: &registry.Info{
				Tags: commonTags,
			},
			expected: func(t *testing.T, u *URL, err error) {
				assert.Nil(t, err)
				assert.NotEmpty(t, u.host)
			},
		},
		{
			desc: "Info with ipv4 addr",
			initAddr: func() net.Addr {
				addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:8888")
				return addr
			},
			info: &registry.Info{
				Tags: commonTags,
			},
			expected: func(t *testing.T, u *URL, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "0.0.0.0:8888", u.host)
			},
		},
		{
			desc: "Info with ipv6 addr",
			initAddr: func() net.Addr {
				addr, _ := net.ResolveTCPAddr("tcp", "[::1]:8888")
				return addr
			},
			info: &registry.Info{
				Tags: commonTags,
			},
			expected: func(t *testing.T, u *URL, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "[::1]:8888", u.host)
			},
		},
		{
			desc:     "Info without params",
			initAddr: commonAddrFunc,
			info: &registry.Info{
				Tags: nil,
			},
			expected: func(t *testing.T, u *URL, err error) {
				assert.Equal(t, errMissingInterface, err)
			},
		},
		{
			desc:     "Info without Interface specified",
			initAddr: commonAddrFunc,
			info: &registry.Info{
				Tags: map[string]string{
					"key": "val",
				},
			},
			expected: func(t *testing.T, u *URL, err error) {
				assert.Equal(t, errMissingInterface, err)
			},
		},
		{
			desc:     "Info with Interface and other information specified",
			initAddr: commonAddrFunc,
			info: &registry.Info{
				Tags: map[string]string{
					DubboServiceInterfaceKey: "interface-val",
					DubboServiceGroupKey:     "g1",
				},
			},
			expected: func(t *testing.T, u *URL, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "interface-val", u.params[dubboInternalInterfaceKey][0])
				assert.Equal(t, "g1", u.params[dubboInternalGroupKey][0])
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			u := new(URL)
			addr := test.initAddr()
			test.info.Addr = addr
			err := u.FromInfo(test.info)
			test.expected(t, u, err)
		})
	}
}
