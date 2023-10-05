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

package main

import (
	"context"

	hessian "github.com/apache/dubbo-go-hessian2"
)

func init() {
	hessian.RegisterPOJO(&Request{})
	hessian.RegisterPOJO(&User{})
	hessian.RegisterPOJO(&ProxyRequest{})
	hessian.RegisterPOJO(&ProxyUser{})
}

type Request struct {
	Name string
}

func (r *Request) JavaClassName() string {
	return "org.apache.dubbo.Request"
}

type ProxyRequest struct {
	Name string
}

func (r *ProxyRequest) JavaClassName() string {
	return "org.apache.dubbo.proxy.Request"
}

type User struct {
	ID   string
	Name string
	Age  int32
}

func (u *User) JavaClassName() string {
	return "org.apache.dubbo.User"
}

type ProxyUser struct {
	ID   string
	Name string
	Age  int32
}

func (u *ProxyUser) JavaClassName() string {
	return "org.apache.dubbo.proxy.User"
}

type UserProvider struct {
	GetUser func(ctx context.Context, req *Request) (*User, error)
}

type UserProviderProxy struct {
	cli *UserProvider
}

func (upp *UserProviderProxy) GetUser(ctx context.Context, req *ProxyRequest) (*ProxyUser, error) {
	userResp, err := upp.cli.GetUser(ctx, &Request{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &ProxyUser{
		ID:   userResp.ID,
		Name: userResp.Name,
		Age:  userResp.Age,
	}, nil
}
