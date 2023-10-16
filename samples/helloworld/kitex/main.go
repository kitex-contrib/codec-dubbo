package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	hello "github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello/greetservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":21000")
	svr := hello.NewServer(new(GreetServiceImpl),
		server.WithServiceAddr(addr),
		server.WithCodec(dubbo.NewDubboCodec(
			dubbo.WithJavaClassName("org.cloudwego.kitex.samples.api.GreetProvider"),
		)),
	)

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
