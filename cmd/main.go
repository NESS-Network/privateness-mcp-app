package main

import (
	"log"
	"os"
	"github.com/jeff-bouchard/privateness-mcp-app/pkg/server"
	"github.com/jeff-bouchard/privateness-mcp-app/pkg/billing"
)

func main() {
	cfg := server.Config{
		Listen:   os.Getenv("LISTEN"),
		BasePath: "/mcp",
		Rates: billing.Rates{PerByteIn: 0.0, PerByteOut: 0.0, PerSecond: 0.0}, // TODO: set real rates
	}
	if err := server.Start(cfg); err != nil { log.Fatal(err) }
}
