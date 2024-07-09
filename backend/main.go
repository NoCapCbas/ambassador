package main

import (
  "log"

  "ambassador/src/database"
  "ambassador/src/routes"

  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
  database.Connect()
  database.AutoMigrate()
  defer database.CloseDB()
  
  // init app
  app := fiber.New()
  // Allow cors and credentials
  app.Use(cors.New(cors.Config{
    AllowOrigins: "http://localhost:8001",
    AllowCredentials: true,
  }))
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
