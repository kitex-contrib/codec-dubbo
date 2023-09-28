package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	"github.com/kitex-contrib/codec-dubbo/tests/benchmark/kitex/server/kitex_gen/user/userservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":20001")
	svr := userservice.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithCodec(dubbo.NewDubboCodec()),
	)

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
