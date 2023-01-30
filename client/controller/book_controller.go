package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamilanindita/go-grpc/client/pb"
)

type BookController interface {
	Route(app *fiber.App, client pb.BookServiceClient)
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
	FindOneById(c *fiber.Ctx) error
}
