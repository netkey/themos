package main

import (
	"flag"
	"log"

	"github.com/netkey/themos/internal/audi/config"

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
