package controllers 

import (
  "github.com/gofiber/fiber/v2"
)

func Hello() error {
  return c.JSON(fiber.Map{
    "message":"hello",
  })
}
