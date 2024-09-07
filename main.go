package main

import (
	"context"
	"log"

	"github.com/jsifuentes/cq-source-jamf/resources/plugin"

	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	if err := serve.Plugin(plugin.Plugin()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
