package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"

	echo "github.com/kitex-contrib/codec-hessian2/tests/kitex/kitex_gen/echo/testservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:20000")
	svr := echo.NewServer(new(TestServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
