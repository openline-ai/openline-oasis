package main

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
	c "openline-ai/oasis-api/config"
	"openline-ai/oasis-api/routes"
)

func main() {
	flag.Parse()
	config := c.Config{}

	if err := env.Parse(&config); err != nil {
		fmt.Printf("missing required config")
		return
	}

	// Our server will live in the routes package
	routes.Run(config.Service.ServerAddress, config)
}
