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

package dubbo2kitex

import (
	"context"
	"helloworld/api"
	"log"
	"net"
	"reflect"
	"testing"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo/testservice"
	"github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo/testsuite"
)

var cli = api.UserProviderClient

// runKitexServer starts Kitex server for testing based on specified address.
// use startCh to tell outer layer that Kitex server has already started.
// use exitCh to receive exiting signal.
func runKitexServer(startCh chan struct{}, exitCh chan error, addr string) {
	netAddr, _ := net.ResolveTCPAddr("tcp", addr)
	svr := testservice.NewServer(
		new(testsuite.TestServiceImpl),
		server.WithServiceAddr(netAddr),
		server.WithExitSignal(func() <-chan error {
			return exitCh
		}),
	)
	server.RegisterStartHook(func() {
		close(startCh)
	})

	if err := svr.Run(); err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	startCh := make(chan struct{})
	exitCh := make(chan error)
	go runKitexServer(startCh, exitCh, "127.0.0.1:20000")
	<-startCh
	// init dubbo cli until kitex srv has started
	if err := config.Load(config.WithPath("./conf/dubbogo.yaml")); err != nil {
		panic(err)
	}
	m.Run()
	exitCh <- nil
}

func TestEchoByte(t *testing.T) {
	var req byte = 12
	resp, err := cli.EchoByte(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoBytes(t *testing.T) {
	req := []byte{1, 2}
	resp, err := cli.EchoBytes(context.Background(), req)
	assertEcho(t, err, req, resp)
}

func TestEchoInt8(t *testing.T) {
	var req int8 = 12
	resp, err := cli.EchoInt8(context.Background(), req)
	assertEcho(t, err, req, resp)
}

// todo(DMwangnima): enhance hessian2.ReflectResponse to support reflecting []int8
//func TestEchoInt8s(t *testing.T) {
//	var req = []int8{1, 2}
//	resp, err := cli.EchoInt8s(context.Background(), req)
//	assertEcho(t, err, req, resp)
//	select {}
//}

func assertEcho(t *testing.T, err error, req, resp interface{}) {
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(req, resp) {
		t.Fatalf("req is not equal to resp, req: %v, resp: %v", req, resp)
	}
}
