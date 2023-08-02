package main

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kitex-contrib/codec-hessian2/tests/kitex/kitex_gen/echo/testservice"
)

func main() {
	cli, err := testservice.NewClient("org.apache.dubbo.UserProvider", client.WithHostPorts("127.0.0.1:20000"))
	if err != nil {
		panic(err)
	}

	resp, err := cli.EchoInt(context.Background(), 0x12345678)
	if err != nil {
		panic(err)
	}
	klog.Infof("resp: %v", resp)
}
