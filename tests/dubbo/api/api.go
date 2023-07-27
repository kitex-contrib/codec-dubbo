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

// User transmission struct; for compatibility with java, field names should be consistent with Java class fields
type User struct {
	ID   string
	Name string
	Age  int32
	Time time.Time
}

func (u *User) JavaClassName() string {
	return "org.apache.dubbo.User" // Should be same as Java class name for java compatibility
}

var UserProviderClient = &UserProvider{} // client pointer

// UserProvider client interface
type UserProvider struct {
	// dubbo tag is necessary to map go function name to java function name
	GetUser func(ctx context.Context, req int32) (*User, error) //`dubbo:"getUser"`
	EchoInt func(ctx context.Context, req int32) (int32, error) //`dubbo:"echoInt"`
}

func init() {
	hessian.RegisterPOJO(&User{}) // Register all transmission struct to hessian lib
	// Register client interface to the framework
	config.SetConsumerService(UserProviderClient)
}
