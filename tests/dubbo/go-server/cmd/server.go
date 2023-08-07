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

package main

import (
	"context"
	"strconv"
	"time"

	"helloworld/api"

	"dubbo.apache.org/dubbo-go/v3/common/logger" // dubbogo 框架日志
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports" // dubbogo 框架依赖，所有dubbogo进程都需要隐式引入一次
)

type UserProvider struct{}

// GetUser implements the interface
func (u *UserProvider) GetUser(ctx context.Context, req int32) (*api.User, error) {
	var err error
	logger.Infof("req:%#v", req)
	user := &api.User{}
	user.ID = strconv.Itoa(int(req))
	user.Name = "laurence"
	user.Age = 22
	user.Time = time.Now()
	return user, err
}

func (u *UserProvider) EchoInt(ctx context.Context, req int32) (int32, error) {
	// for exception test
	// return 0, errors.New("EchoInt failed without reason")

	return req, nil
}

// MethodMapper is for mapping go func name to java func name.
// Not necessary for go client -> go server
// func (s *UserProvider) MethodMapper() map[string]string {
// 	return map[string]string{
// 		"GetUser": "getUser",
// 	}
// }

func init() {
	config.SetProviderService(&UserProvider{}) // Register service provider, should be same in the config file
}

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}
	select {}
}
