package main

import (
	"context"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	hello "github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello/greetservice"

	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	serviceName := "helloworld"

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	addr, _ := net.ResolveTCPAddr("tcp", ":21000")
	svr := hello.NewServer(new(GreetServiceImpl),
		server.WithServiceAddr(addr),
		server.WithCodec(dubbo.NewDubboCodec(
			dubbo.WithJavaClassName("org.cloudwego.kitex.samples.api.GreetProvider"),
		)),
		server.WithSuite(tracing.NewServerSuite()),
	)

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
