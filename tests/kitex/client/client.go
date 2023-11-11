package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"

	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo"
	"github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo/testservice"
)

func main() {
	cli, err := testservice.NewClient("test",
		//client.WithHostPorts("127.0.0.1:20000"),
		client.WithHostPorts("192.168.1.26:20001"),
		client.WithCodec(dubbo.NewDubboCodec(
			dubbo.WithJavaClassName("org.apache.dubbo.tests.api.UserProvider"),
			dubbo.WithFileDescriptor(echo.GetFileDescriptorForApi()),
		)),
	)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	//resp, err := cli.EchoInt(context.Background(), 0x12345678)
	//if err != nil {
	//	panic(err)
	//}
	//klog.Infof("resp: %v", resp)
	//
	//_, err = cli.EchoInt(context.Background(), 400)
	//if err == nil {
	//	panic("want err but got nothing")
	//}
	//klog.Infof("got err: %v", err)

	resA, err := cli.EchoMethodA(ctx, true)
	check(err)
	klog.Infof("resA: %v", resA)

	resB, err := cli.EchoMethodB(ctx, 1)
	check(err)
	klog.Infof("resA: %v", resB)

	resC, err := cli.EchoMethodC(ctx, 1)
	check(err)
	klog.Infof("resA: %v", resC)

	resD, err := cli.EchoMethodD(ctx, true, 1)
	check(err)
	klog.Infof("resA: %v", resD)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
