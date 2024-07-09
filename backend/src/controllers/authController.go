package controllers

import (
  "github.com/gofiber/fiber/v2"
  "ambassador/src/models"
  "ambassador/src/database"
  "ambassador/src/utils"
  "log"
  "fmt"
)

func Register(c *fiber.Ctx) error {
  var data map[string]string

  if err := c.BodyParser(&data); err != nil {
    return err
  }

  if data["password"] != data["password_confirm"] {
    c.Status(400)
    return c.JSON(fiber.Map{
      "message":"passwords do not match",
    })
  }

  password, _ := utils.HashPassword(data["password"])

  user := models.User{
    FirstName: data["first_name"],
    LastName: data["last_name"],
    Email: data["email"],
    Password: password,
    IsAmbassador: false,
  }

  // insert user into database
  err := database.CreateUser(&user)
  if err != nil {
    log.Printf("Failed to create user: %v\n", err)
    c.Status(fiber.StatusBadRequest)
    return c.JSON(fiber.Map{
      "message":"Could not create user",
    })
  }

  return c.JSON(user)

}

func Login(c *fiber.Ctx) error {
  var data map[string]string 
  var err error

  if err = c.BodyParser(&data); err != nil {
   return err
  }
  
  var user *models.User
  user, err = database.AuthenticateUser(data["email"], data["password"])
  if err != nil {
    log.Printf("Failed to authenticate user: %v\n", err)
    c.Status(fiber.StatusBadRequest)
    return fmt.Errorf("Error Authenticating User: %v", err)
  }

  return c.JSON(user)
}
