package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kamilanindita/go-grpc/client/config"
	"github.com/kamilanindita/go-grpc/client/controller"
	"github.com/kamilanindita/go-grpc/client/exception"
	"github.com/kamilanindita/go-grpc/client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// Setup config
	configuration := config.New()

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	conn, err := grpc.Dial(configuration.Get("GRPC_SERVER")+":"+configuration.Get("GRPC_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		exception.PanicIfNeeded(err)
	}
	defer conn.Close()

	bookServiceClient := pb.NewBookServiceClient(conn)

	// Setup Routing
	controller.BookServiceRoute(app, bookServiceClient)

	// Start App
	app.Listen(":" + configuration.Get("PORT"))

}
