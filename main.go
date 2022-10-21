package main

import (
	"flag"
	"openline-ai/openline-oasis/routes"
)

var (
	addr = flag.String("addr", ":8006", "The server address")
)

func main() {
	flag.Parse()
	// Our server will live in the routes package
	routes.Run(*addr)
}
