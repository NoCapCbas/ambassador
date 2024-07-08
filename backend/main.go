package main

import (
  "log"

  "ambassador/src/database"
  "ambassador/src/routes"

  "github.com/gofiber/fiber/v2"
)

func main() {
  database.Connect()
  database.AutoMigrate()
  defer database.CloseDB()
  
  // init app
  app := fiber.New()
  // set up routes
  routes.Setup(app)
  // example root route
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello World")
  })

  // start server at port
  err := app.Listen(":8000")
  if err != nil {
    log.Fatalf("Error starting server: %v\n", err)
  }
}
