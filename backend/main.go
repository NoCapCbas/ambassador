package main

import (
  "ambassador/src/database"
  "ambassador/src/routes"

  "github.com/gofiber/fiber/v2"
)

func main() {
  database.Connect()
  database.AutoMigrate()
  
  // init app
  app := fiber.New()
  // handler
  routes.Setup(app)
  // start server at port
  app.Listen(":8000")
}
