package main

import (
	"barry/config"
	"barry/router"
	"fmt"
)

//go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64
//go env -w CGO_ENABLED=1 GOOS=windows GOARCH=amd64

func main() {
	routers := router.InitRouter()
	startErr := routers.Run(config.Addr)
	if startErr != nil {
		fmt.Println("start proj error")
	}
}