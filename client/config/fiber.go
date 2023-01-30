package config

import (
	"github.com/kamilanindita/go-grpc/client/exception"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
