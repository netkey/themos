package main

import (
	"flag"
	"log"

	"github.com/netkey/audi/config"
)

func main() {
	log.Println("Starting microservice")

	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

}
