package main

import (
	"context"

	"github.com/PandeKaustubhS/microservice-todo/config"
	server "github.com/PandeKaustubhS/microservice-todo/server"
)

func init() {
	config.Load()
}

func main() {
	conf := config.Gateway()
	ctx := context.Background()
	server.StartRESTGateway(ctx, conf)
}
