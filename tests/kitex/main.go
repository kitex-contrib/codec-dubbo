package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"

	echo "github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo/testservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":20000")
	svr := echo.NewServer(new(TestServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
