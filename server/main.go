package main

import (
	"github.com/kamilanindita/go-grpc/server/config"
	"github.com/kamilanindita/go-grpc/server/delivery/grpc"
)

func main() {
	// Setup config
	configuration := config.New()
	database := config.NewDatabase(configuration)

	// Init grpc
	grpc.Init(configuration, database)

}
