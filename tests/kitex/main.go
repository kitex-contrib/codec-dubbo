package main

import (
	"log"

	echo "github.com/kitex-contrib/codec-hessian2/tests/kitex/kitex_gen/echo/testservice"
)

func main() {
	svr := echo.NewServer(new(TestServiceImpl))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
