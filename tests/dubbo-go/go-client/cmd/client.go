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

	"github.com/dubbogo/gost/log/logger"

	"helloworld/api"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

func main() {
	// 启动框架
	if err := config.Load(); err != nil {
		panic(err)
	}

	// 发起调用
	var i int32 = 0x0A0B0C0D
	/*
		user, err := api.UserProviderClient.GetUser(context.TODO(), i)
		if err != nil {
			panic(err)
		}
		logger.Infof("response result: %+v", user)
	*/

	resp, err := api.UserProviderClient.EchoInt(context.TODO(), i)
	if err != nil {
		panic(err)
	}
	logger.Infof("response result: %+v", resp)

	byteResp, err := api.UserProviderClient.EchoByte(context.TODO(), 8)
	if err != nil {
		panic(err)
	}
	logger.Infof("response result: %+v", byteResp)

	_, err = api.UserProviderClient.EchoInt(context.TODO(), 400)
	if err == nil {
		panic("want err but got nothing")
	}
	logger.Infof("got err: %+v", err)
}
