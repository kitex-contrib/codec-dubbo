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

package registry_test

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"fmt"
	"github.com/cloudwego/kitex/client"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	"github.com/kitex-contrib/codec-dubbo/registries/zookeeper/resolver"
	"github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo/testservice"
	"github.com/stretchr/testify/assert"
	"helloworld/api"
	"net"
	"os/exec"
	"testing"
	"time"
)

func runDubboGoServer(exitChan chan struct{}) {
	config.SetProviderService(&api.UserProviderImpl{})
	// multiple version implementation with same Interface
	config.SetProviderService(&api.UserProviderImplV1{})
	if err := config.Load(config.WithPath("./conf/dubbogo.yaml")); err != nil {
		panic(err)
	}
	select {
	case <-exitChan:
		return
	}
}

func runDubboJavaServer() context.CancelFunc {
	testDir := "../dubbo-java"
	// initialize mvn packages
	cleanCmd := exec.Command("mvn", "clean", "package")
	cleanCmd.Dir = testDir
	if _, err := cleanCmd.Output(); err != nil {
		panic(fmt.Sprintf("mvn clean package failed: %s", err))
	}
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "mvn",
		"-Djava.net.preferIPv4Stack=true",
		"-Dexec.mainClass=org.apache.dubbo.tests.provider.RegistryApplication",
		"exec:java")
	cmd.Dir = testDir

	go func() {
		if err := cmd.Run(); err != nil {
			panic(fmt.Sprintf("mvn exec failed: %s", err))
		}
	}()

	return cancel
}

func waitForPort(port string) {
	for {
		conn, err := net.Dial("tcp", net.JoinHostPort("127.0.0.1", port))
		if err != nil {
			time.Sleep(time.Second * 1)
		} else {
			conn.Close()
			return
		}
	}
}

func TestMain(m *testing.M) {
	exitChan := make(chan struct{})
	go runDubboGoServer(exitChan)
	cancel := runDubboJavaServer()
	waitForPort("20000")
	waitForPort("20001")
	m.Run()
	// close dubbo-go server
	close(exitChan)
	// kill dubbo-java server
	cancel()
}

func TestResolve(t *testing.T) {
	// please refer to ./conf/dubbogo.yaml
	goInterfaceName := "org.apache.dubbo.tests.go.api.UserProvider"
	javaInterfaceName := "org.apache.dubbo.tests.api.UserProvider"
	tests := []struct {
		resOpts   []resolver.Option
		codecOpts []dubbo.Option
		judge     func(t *testing.T, cli testservice.Client)
	}{
		{
			resOpts: []resolver.Option{
				resolver.WithServers("127.0.0.1:2181"),
				resolver.WithInterfaceName(goInterfaceName),
				resolver.WithRegistryGroup("myGroup"),
			},
			codecOpts: []dubbo.Option{
				dubbo.WithJavaClassName(goInterfaceName),
			},
			judge: func(t *testing.T, cli testservice.Client) {
				resp, err := cli.EchoBool(context.Background(), true)
				assert.Nil(t, err)
				assert.Equal(t, true, resp)
			},
		},
		{
			resOpts: []resolver.Option{
				resolver.WithServers("127.0.0.1:2181"),
				resolver.WithInterfaceName(goInterfaceName),
				resolver.WithRegistryGroup("myGroup"),
				resolver.WithServiceGroup("g1"),
				resolver.WithServiceVersion("v1"),
			},
			codecOpts: []dubbo.Option{
				dubbo.WithJavaClassName(goInterfaceName),
			},
			judge: func(t *testing.T, cli testservice.Client) {
				resp, err := cli.EchoBool(context.Background(), true)
				assert.Nil(t, err)
				assert.Equal(t, false, resp)
			},
		},
		{
			resOpts: []resolver.Option{
				resolver.WithServers("127.0.0.1:2181"),
				resolver.WithInterfaceName(javaInterfaceName),
				resolver.WithRegistryGroup("myGroup"),
			},
			codecOpts: []dubbo.Option{
				dubbo.WithJavaClassName(javaInterfaceName),
			},
			judge: func(t *testing.T, cli testservice.Client) {
				resp, err := cli.EchoBool(context.Background(), true)
				assert.Nil(t, err)
				assert.Equal(t, true, resp)
			},
		},
		{
			resOpts: []resolver.Option{
				resolver.WithServers("127.0.0.1:2181"),
				resolver.WithInterfaceName(javaInterfaceName),
				resolver.WithRegistryGroup("myGroup"),
				resolver.WithServiceGroup("g1"),
				resolver.WithServiceVersion("v1"),
			},
			codecOpts: []dubbo.Option{
				dubbo.WithJavaClassName(javaInterfaceName),
			},
			judge: func(t *testing.T, cli testservice.Client) {
				resp, err := cli.EchoBool(context.Background(), true)
				assert.Nil(t, err)
				assert.Equal(t, false, resp)
			},
		},
	}

	for _, test := range tests {
		res, err := resolver.NewZookeeperResolver(test.resOpts...)
		assert.Nil(t, err)
		cli, err := testservice.NewClient("testtest",
			client.WithResolver(res),
			client.WithCodec(dubbo.NewDubboCodec(test.codecOpts...)),
		)
		assert.Nil(t, err)
		test.judge(t, cli)
	}
}
