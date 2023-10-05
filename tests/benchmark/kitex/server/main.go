package main

import (
	"flag"
	"log"
	"net"
	"strconv"

	"github.com/cloudwego/kitex/server"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	"github.com/kitex-contrib/codec-dubbo/tests/benchmark/kitex/server/kitex_gen/user/userservice"
)

func main() {
	var srvPort int
	flag.IntVar(&srvPort, "p", 20001, "")
	flag.Parse()
	addr, _ := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(srvPort))
	svr := userservice.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithCodec(dubbo.NewDubboCodec()),
	)

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
