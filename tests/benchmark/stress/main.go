package main

import (
	"context"
	"flag"

	"dubbo.apache.org/dubbo-go/v3/config"

	"github.com/dubbogo/tools/pkg/stressTest"
)

func main() {
	cli := new(UserProvider)
	config.SetConsumerService(cli)
	if err := config.Load(config.WithPath("./dubbogo.yml")); err != nil {
		panic(err)
	}

	var tps, parallel, payloadLen int
	flag.IntVar(&tps, "t", 0, "")
	flag.IntVar(&parallel, "p", 0, "")
	flag.IntVar(&payloadLen, "l", 0, "")
	flag.Parse()

	ctx := context.Background()
	req := "laurence" + string(make([]byte, payloadLen))
	stressTest.NewStressTestConfigBuilder().
		SetTPS(tps).
		SetParallel(parallel).
		SetDuration("1m").
		Build().
		Start(func() error {
			if _, err := cli.GetUser(ctx, &Request{
				Name: req,
			}); err != nil {
				return err
			}
			return nil
		})
}
