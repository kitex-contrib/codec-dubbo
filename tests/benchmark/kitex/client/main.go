package main

import (
	"flag"
	"log"
	"net"
	"strconv"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	"github.com/kitex-contrib/codec-dubbo/tests/benchmark/kitex/client/kitex_gen/user/proxyservice"
	"github.com/kitex-contrib/codec-dubbo/tests/benchmark/kitex/server/kitex_gen/user/userservice"
)

func main() {
	var cliPort int
	var srvAddr string
	flag.IntVar(&cliPort, "p", 20000, "")
	flag.StringVar(&srvAddr, "addr", "127.0.0.1:20001", "")
	flag.Parse()

	cliCodec := dubbo.NewDubboCodec(
		dubbo.WithJavaClassName("org.apache.dubbo.UserProvider"),
	)
	cli, err := userservice.NewClient("test",
		client.WithHostPorts(srvAddr),
		client.WithCodec(cliCodec),
	)
	if err != nil {
		panic(err)
	}
	srvCodec := dubbo.NewDubboCodec(
		dubbo.WithJavaClassName("org.apache.dubbo.UserProviderProxy"),
	)
	addr, _ := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(cliPort))
	svr := proxyservice.NewServer(&ProxyServiceImpl{cli: cli},
		server.WithServiceAddr(addr),
		server.WithCodec(srvCodec),
	)

	if err = svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
