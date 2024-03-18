package main

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	"github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello"
	"github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello/greetservice"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	serviceName := "helloworld-client"
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	cli, err := greetservice.NewClient("helloworld",
		client.WithHostPorts("127.0.0.1:21000"),
		client.WithCodec(
			dubbo.NewDubboCodec(
				dubbo.WithJavaClassName("org.cloudwego.kitex.samples.api.GreetProvider"),
			),
		),
		client.WithSuite(tracing.NewClientSuite()),
	)
	if err != nil {
		panic(err)
	}

	resp, err := cli.Greet(context.Background(), "world")
	if err != nil {
		klog.Error(err)
		return
	}
	klog.Infof("resp: %s", resp)

	respWithStruct, err := cli.GreetWithStruct(context.Background(), &hello.GreetRequest{Req: "world"})
	if err != nil {
		klog.Error(err)
		return
	}
	klog.Infof("respWithStruct: %s", respWithStruct.Resp)
}
