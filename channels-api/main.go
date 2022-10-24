package main

import (
	"flag"
	"openline-ai/openline-channels/routes"
)

var (
	addr = flag.String("addr", ":8013", "The server address")
)

func main() {
	flag.Parse()
	// Our server will live in the routes package
	routes.Run(*addr)
}
