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
	"flag"
	"strconv"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

func main() {
	var cliPort int
	var srvAddr string
	flag.IntVar(&cliPort, "p", 20000, "")
	flag.StringVar(&srvAddr, "addr", "127.0.0.1:20001", "")
	flag.Parse()

	cli := new(UserProvider)
	config.SetConsumerService(cli)
	config.SetProviderService(&UserProviderProxy{cli: cli})
	protCfg := config.NewProtocolConfigBuilder().
		SetName("dubbo").
		SetPort(strconv.Itoa(cliPort)).
		Build()
	refCfg := config.NewReferenceConfigBuilder().
		SetProtocol("dubbo").
		SetInterface("org.apache.dubbo.UserProvider").
		SetURL("dubbo://" + srvAddr).
		Build()
	conCfg := config.NewConsumerConfigBuilder().
		AddReference("UserProvider", refCfg).
		Build()
	svcCfg := config.NewServiceConfigBuilder().
		SetProtocolIDs("dubbo").
		SetInterface("org.apache.dubbo.UserProviderProxy").
		Build()
	proCfg := config.NewProviderConfigBuilder().
		AddService("UserProviderProxy", svcCfg).
		Build()
	logCfg := config.NewLoggerConfigBuilder().
		SetLevel("error").
		Build()
	rootCfg := config.NewRootConfigBuilder().
		AddProtocol("dubbo", protCfg).
		SetConsumer(conCfg).
		SetProvider(proCfg).
		SetLogger(logCfg).
		Build()
	if err := rootCfg.Init(); err != nil {
		panic(err)
	}
	select {}
}
