package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	"github.com/kitex-contrib/codec-dubbo/tests/benchmark/kitex/client/kitex_gen/user/benchmarkservice"
	"github.com/kitex-contrib/codec-dubbo/tests/benchmark/kitex/server/kitex_gen/user/userservice"
)

func main() {
	codec := dubbo.NewDubboCodec()
	cli, err := userservice.NewClient("test",
		client.WithHostPorts("127.0.0.1:20001"),
		client.WithCodec(codec),
	)
	if err != nil {
		panic(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", ":20000")
	svr := benchmarkservice.NewServer(&BenchmarkServiceImpl{cli: cli},
		server.WithServiceAddr(addr),
		server.WithCodec(codec),
	)

	if err = svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
