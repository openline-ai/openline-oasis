package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	c "openline-ai/channels-api/config"
	"openline-ai/channels-api/routes"
)

func main() {
	conf := c.Config{}
	if err := env.Parse(&conf); err != nil {
		fmt.Printf("missing required config")
		return
	}
	// Our server will live in the routes package
	routes.Run(conf)
}
