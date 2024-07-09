package controllers

import (
  "github.com/gofiber/fiber/v2"
  "github.com/dgrijalva/jwt-go"
  "ambassador/src/models"
  "ambassador/src/database"
  "strconv"
  "log"
  "fmt"
  "time"
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

  user := models.User{
    FirstName: data["first_name"],
    LastName: data["last_name"],
    Email: data["email"],
    IsAmbassador: false,
  }

  user.SetPassword(data["password"])


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
    c.Status(fiber.StatusBadRequest)
    return c.JSON(fiber.Map{
      "message":"Cannot parse JSON",
    }) 
  }
  
  var user *models.User
  user, err = database.AuthenticateUser(data["email"], data["password"])
  if err != nil {
    log.Printf("Failed to authenticate user: %v\n", err)
    c.Status(fiber.StatusBadRequest)
    return fmt.Errorf("Invalid email or password", err)
  }

  payload := jwt.StandardClaims{
    Subject: strconv.Itoa(int(user.ID)),
    ExpiresAt: time.Now().Add(time.Hour*24).Unix(),
  }

  token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))
  if err != nil {
    log.Printf("Failed to generate JWT token: %v\n", err)
    c.Status(fiber.StatusBadRequest)
    return c.JSON(fiber.Map{
      "message":"Invalid Token Credentials",
    })
  }

  cookie := fiber.Cookie{
    Name: "jwt",
    Value: token,
    Expires: time.Now().Add(time.Hour*24),
    HTTPOnly: true,
  }
  
  c.Cookie(&cookie)

  return c.JSON(fiber.Map{
    "message":"success",
  })
}
