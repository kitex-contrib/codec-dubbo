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
}

type Request struct {
	Name string
}

func (r *Request) JavaClassName() string {
	return "org.apache.dubbo.Request"
}

type User struct {
	ID   string
	Name string
	Age  int32
}

func (u *User) JavaClassName() string {
	return "org.apache.dubbo.User"
}

type UserProvider struct{}

func (up *UserProvider) GetUser(ctx context.Context, req *Request) (*User, error) {
	return &User{
		ID:   "12345",
		Name: "Hello " + req.Name,
		Age:  21,
	}, nil
}
