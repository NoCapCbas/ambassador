package main

import (
  "ambassador/src/database"

  "github.com/gofiber/fiber/v2"
)

func main() {
  database.Connect()
  database.AutoMigrate()
  
  // init app
  app := fiber.New()
  // handler
  app.Get("/", func (c *fiber.Ctx) error {
      return c.SendString("Hello, Test!")
  })
  // start server at port
  app.Listen(":8000")
}
