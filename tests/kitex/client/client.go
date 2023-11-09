package main

import (
	"context"

	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo"
	"github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo/testservice"
)

func main() {
	cli, err := testservice.NewClient("test",
		client.WithHostPorts("127.0.0.1:20000"),
		client.WithCodec(dubbo.NewDubboCodec(
			dubbo.WithJavaClassName("org.apache.dubbo.tests.api.UserProvider"),
			dubbo.WithFileDescriptor(echo.GetFileDescriptorForApi()),
		)),
	)
	if err != nil {
		panic(err)
	}

	resp, err := cli.EchoInt(context.Background(), 0x12345678)
	if err != nil {
		panic(err)
	}
	klog.Infof("resp: %v", resp)

	_, err = cli.EchoInt(context.Background(), 400)
	if err == nil {
		panic("want err but got nothing")
	}
	klog.Infof("got err: %v", err)
}
