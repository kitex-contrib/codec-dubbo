package main

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

func main() {
	config.SetProviderService(&UserProvider{})
	if err := config.Load(config.WithPath("./dubbogo.yml")); err != nil {
		panic(err)
	}
	select {}
}
