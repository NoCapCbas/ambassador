package models

type User struct { 
  ID uint `json:"id"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  Email string `json:"email"`
  Password []byte `json:"password"`
  IsAmbassador bool `json:"is_ambassador"`
}
