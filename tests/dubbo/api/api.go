/*
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

package api

import (
	"context"
	"time"

	"dubbo.apache.org/dubbo-go/v3/config"
	hessian "github.com/apache/dubbo-go-hessian2"
)

// 1. 定义传输结构， 如需 Java 互通，字段需要与 Java 侧对应，首字母大写
type User struct {
	ID   string
	Name string
	Age  int32
	Time time.Time
}

func (u *User) JavaClassName() string {
	return "org.apache.dubbo.User" // 如果与 Java 互通，需要与 Java 侧 User class全名对应,
}

var UserProviderClient = &UserProvider{} // 客户端指针

// 2。 定义客户端存根类：UserProvider
type UserProvider struct {
	// dubbo标签，用于适配go侧客户端大写方法名 -> java侧小写方法名，只有 dubbo 协议客户端才需要使用
	GetUser func(ctx context.Context, req int32) (*User, error) //`dubbo:"getUser"`
	EchoInt func(ctx context.Context, req int32) (int32, error) //`dubbo:"echoInt"`
}

func init() {
	hessian.RegisterPOJO(&User{}) // 注册传输结构到 hessian 库
	// 注册客户端存根类到框架，实例化客户端接口指针 userProvider
	config.SetConsumerService(UserProviderClient)
}
