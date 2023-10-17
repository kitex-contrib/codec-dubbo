package main

import (
	"log"
	"net"

	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"

	"github.com/cloudwego/kitex/server"

	echo "github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo/testservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":20000")
	svr := echo.NewServer(new(TestServiceImpl),
		server.WithServiceAddr(addr),
		server.WithCodec(dubbo.NewDubboCodec(
			dubbo.WithJavaClassName("org.apache.dubbo.tests.api.UserProvider"),
		)),
	)

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
