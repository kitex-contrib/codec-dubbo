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

package kitex2dubbo

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/cloudwego/kitex/client"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	"github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo/testservice"
	"helloworld/api"
	"reflect"
	"testing"
	"time"
)

var cli2Go, cli2Java testservice.Client

// runDubboServer starts dubbo-go server.
// use exitChan to receive exit signal.
func runDubboGoServer(exitChan chan struct{}) {
	config.SetProviderService(&api.UserProviderImpl{})
	if err := config.Load(config.WithPath("./conf/dubbogo.yaml")); err != nil {
		panic(err)
	}
	select {
	case <-exitChan:
		return
	}
}

func TestMain(m *testing.M) {
	exitChan := make(chan struct{})
	go runDubboGoServer(exitChan)
	cancel := runDubboJavaServer()
	//wait for dubbo-go and dubbo-java server initialization
	time.Sleep(10 * time.Second)
	var err error
	cli2Go, err = testservice.NewClient("test",
		client.WithHostPorts("127.0.0.1:20000"),
		client.WithCodec(dubbo.NewDubboCodec()),
	)
	if err != nil {
		panic(err)
	}
	cli2Java, err = testservice.NewClient("test",
		client.WithHostPorts("127.0.0.1:20001"),
		client.WithCodec(dubbo.NewDubboCodec()),
	)
	if err != nil {
		panic(err)
	}
	m.Run()
	// close dubbo-go server
	close(exitChan)
	// kill dubbo-java server
	cancel()
}

func TestEchoBool(t *testing.T) {
	req := true
	resp, err := cli2Go.EchoBool(context.Background(), req)
	assertEcho(t, err, req, resp)
}

// dubbo-go does not support
//func TestEchoByte(t *testing.T) {
//	var req int8 = 12
//	resp, err := cli2Go.EchoByte(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

// dubbo-go does not support
//func TestEchoInt16(t *testing.T) {
//	var req int16 = 12
//	resp, err := cli2Go.EchoInt16(context.Background(), req)
//	assertEcho(t, err, req, resp)
//}

func TestEchoInt32(t *testing.T) {
	var req int32 = 12
	resp, err := cli2Go.EchoInt32(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoInt64(t *testing.T) {
	var req int64 = 12
	resp, err := cli2Go.EchoInt64(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoDouble(t *testing.T) {
	var req float64 = 12.3456
	resp, err := cli2Go.EchoDouble(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoString(t *testing.T) {
	req := "12"
	resp, err := cli2Go.EchoString(context.Background(), req)
	assertEcho(t, err, req, resp)
}

// ----------kitex -> dubbo-java----------

func TestEchoBool_Java(t *testing.T) {
	req := true
	resp, err := cli2Java.EchoBool(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoByte_Java(t *testing.T) {
	var req int8 = 12
	resp, err := cli2Java.EchoByte(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoInt16_Java(t *testing.T) {
	var req int16 = 12
	resp, err := cli2Java.EchoInt16(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoInt32_Java(t *testing.T) {
	var req int32 = 12
	resp, err := cli2Java.EchoInt32(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoInt64_Java(t *testing.T) {
	var req int64 = 12
	resp, err := cli2Java.EchoInt64(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoDouble_Java(t *testing.T) {
	var req float64 = 12.3456
	resp, err := cli2Java.EchoDouble(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoString_Java(t *testing.T) {
	req := "12"
	resp, err := cli2Java.EchoString(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func assertEcho(t *testing.T, err error, req, resp interface{}) {
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(req, resp) {
		t.Fatalf("req is not equal to resp, req: %v, resp: %v", req, resp)
	}
}
