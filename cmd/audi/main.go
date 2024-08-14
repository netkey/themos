package main

import (
	"flag"

	"github.com/netkey/themos/internal/audi/config"
	"github.com/netkey/themos/pkg/log"

	_ "go.uber.org/automaxprocs"
)

func main() {
	log.Printf("Starting microservice: %s", config.ServiceName)

	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

}
