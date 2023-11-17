package registry

import (
	"context"
	"helloworld/api"
	"net"
	"testing"
	"time"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/cloudwego/kitex/client"
	registry2 "github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	"github.com/kitex-contrib/codec-dubbo/registries"
	"github.com/kitex-contrib/codec-dubbo/registries/zookeeper/registry"
	"github.com/kitex-contrib/codec-dubbo/registries/zookeeper/resolver"
	"github.com/kitex-contrib/codec-dubbo/tests/kitex/handler"
	"github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo/testservice"
	"github.com/kitex-contrib/codec-dubbo/tests/util"
	"github.com/stretchr/testify/assert"
)

func TestRegistry(t *testing.T) {
	zookeeperAddress1 := "127.0.0.1:2181"
	kitexAddress := ":20000"
	javaInterfaceName := "org.apache.dubbo.tests.api.UserProvider"
	tests := []struct {
		regOpts      []registry.Option
		codecOpts    []dubbo.Option
		tags         map[string]string
		cliResOpts   []resolver.Option
		cliCodecOpts []dubbo.Option
		cliOpts      []client.Option
		cliCfgPath   string
	}{
		{
			regOpts: []registry.Option{
				registry.WithServers(zookeeperAddress1),
				registry.WithRegistryGroup("myGroup"),
			},
			codecOpts: []dubbo.Option{
				dubbo.WithJavaClassName(javaInterfaceName),
			},
			tags: map[string]string{
				registries.DubboServiceInterfaceKey:   javaInterfaceName,
				registries.DubboServiceApplicationKey: "dubbo",
			},
			cliResOpts: []resolver.Option{
				resolver.WithServers(zookeeperAddress1),
				resolver.WithRegistryGroup("myGroup"),
			},
			cliCodecOpts: []dubbo.Option{
				dubbo.WithJavaClassName(javaInterfaceName),
			},
			cliOpts: []client.Option{
				client.WithTag(registries.DubboServiceInterfaceKey, javaInterfaceName),
			},
			cliCfgPath: "./conf/dubbogo.yaml",
		},
	}

	for _, test := range tests {
		startChan := make(chan struct{})
		exitChan := make(chan error)
		exitFinish := make(chan struct{})
		go startKitexServerWithRegistry(t, startChan, exitChan, exitFinish, kitexAddress, test.regOpts, test.codecOpts, test.tags)
		<-startChan
		// wait for registering
		time.Sleep(3 * time.Second)
		testKitexClient(t, test.cliResOpts, test.cliCodecOpts, test.cliOpts)
		testDubboJavaClient(t)
		testDubboGoClient(t, test.cliCfgPath)
		exitChan <- nil
		<-exitFinish
	}
}

func startKitexServerWithRegistry(t *testing.T, startCh chan struct{}, exitCh chan error, exitFinish chan struct{},
	addr string, regOpts []registry.Option, codecOpts []dubbo.Option, tags map[string]string,
) {
	netAddr, _ := net.ResolveTCPAddr("tcp", addr)
	reg, err := registry.NewZookeeperRegistry(regOpts...)
	assert.Nil(t, err)
	svr := testservice.NewServer(
		new(handler.TestServiceImpl),
		server.WithServiceAddr(netAddr),
		server.WithRegistry(reg),
		server.WithRegistryInfo(&registry2.Info{
			Tags: tags,
		}),
		server.WithCodec(dubbo.NewDubboCodec(codecOpts...)),
		server.WithExitSignal(func() <-chan error {
			return exitCh
		}),
	)
	server.RegisterStartHook(func() {
		close(startCh)
	})
	assert.Nil(t, svr.Run())
	exitFinish <- struct{}{}
}

func testKitexClient(t *testing.T, resOpts []resolver.Option, codecOpts []dubbo.Option, cliOpts []client.Option) {
	res, err := resolver.NewZookeeperResolver(resOpts...)
	assert.Nil(t, err)
	opts := []client.Option{
		client.WithResolver(res),
		client.WithCodec(
			dubbo.NewDubboCodec(codecOpts...),
		),
	}
	opts = append(opts, cliOpts...)
	cli, err := testservice.NewClient("testtest", opts...)
	assert.Nil(t, err)
	resp, err := cli.EchoBool(context.Background(), true)
	assert.Nil(t, err)
	assert.True(t, resp)
}

func testDubboGoClient(t *testing.T, configPath string) {
	cli := api.UserProviderClient
	err := config.Load(config.WithPath(configPath))
	assert.Nil(t, err)
	resp, err := cli.EchoBool(context.Background(), true)
	assert.Nil(t, err)
	assert.True(t, resp)
}

func testDubboJavaClient(t *testing.T) {
	util.RunAndTestDubboJavaClient(t, "../../dubbo-java", "org.apache.dubbo.tests.client.Application",
		[]string{"withRegistry"},
		[]string{
			"EchoBool",
		},
	)
}
