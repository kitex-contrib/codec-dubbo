package main

import (
	"log"
	"net"

	"github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo"

	"github.com/cloudwego/kitex/server"

	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	"github.com/kitex-contrib/codec-dubbo/tests/kitex/handler"
	"github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo/testservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":20000")
	svr := testservice.NewServer(new(handler.TestServiceImpl),
		server.WithServiceAddr(addr),
		server.WithCodec(dubbo.NewDubboCodec(
			dubbo.WithJavaClassName("org.apache.dubbo.tests.api.UserProvider"),
			dubbo.WithFileDescriptor(echo.GetFileDescriptorForApi()),
		)),
	)

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
