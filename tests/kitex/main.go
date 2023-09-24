package main

import (
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"

	echo "github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo/testservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":20000")
	svr := echo.NewServer(new(TestServiceImpl),
		server.WithServiceAddr(addr),
		server.WithCodec(dubbo.NewDubboCodec()),
	)

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
