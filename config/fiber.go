package config

import (
	"auth-service/app/common/exceptions"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		ErrorHandler: exceptions.ErrorHandler,
	}
}
