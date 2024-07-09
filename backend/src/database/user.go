package database

import (
  "database/sql"
  "fmt"
  "ambassador/src/models"
  "ambassador/src/utils"
)
// CreateUser
func CreateUser(user *models.User) error {
  query := `
    INSERT INTO users (first_name, last_name, email, password, is_ambassador)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id
  `

  err := db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password, user.IsAmbassador)
  if err != nil {
    return fmt.Errorf("CreateUser: %v", err)
  }
  return nil
}

// GetUser retrieves a user from the database by ID.
func GetUser(id int) (models.User, error) {
    var user models.User
    query := `
        SELECT id, first_name, last_name, email, password, is_ambassador
        FROM users
        WHERE id = $1
    `
    err := db.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.IsAmbassador)
    if err != nil {
        if err == sql.ErrNoRows {
            return user, fmt.Errorf("GetUser: user with id %d not found", id)
        }
        return user, fmt.Errorf("GetUser: %v", err)
    }
    return user, nil
}

// UpdateUser updates an existing user in the database.
func UpdateUser(user models.User) error {
    query := `
        UPDATE users
        SET first_name = $1, last_name = $2, email = $3, password = $4, is_ambassador = $5
        WHERE id = $6
    `
    _, err := db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password, user.IsAmbassador, user.ID)
    if err != nil {
        return fmt.Errorf("UpdateUser: %v", err)
    }
    return nil
}

// DeleteUser deletes a user from the database by ID.
func DeleteUser(id int) error {
    query := `
        DELETE FROM users
        WHERE id = $1
    `
    _, err := db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("DeleteUser: %v", err)
    }
    return nil
}

// GetAllUsers retrieves all users from the database.
func GetAllUsers() ([]models.User, error) {
    query := `
        SELECT id, first_name, last_name, email, password, is_ambassador
        FROM users
    `
    rows, err := db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("GetAllUsers: %v", err)
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.IsAmbassador); err != nil {
            return nil, fmt.Errorf("GetAllUsers: %v", err)
        }
        users = append(users, user)
    }
    return users, nil
}

// AuthenticateUser authenticates a user by email and password.
func AuthenticateUser(email, password string) (*models.User, error) {
    var user models.User
    query := `
        SELECT id, first_name, last_name, email, password, is_ambassador
        FROM users
        WHERE email = $1
    `
    err := db.QueryRow(query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.IsAmbassador)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("AuthenticateUser: user not found")
        }
        return nil, fmt.Errorf("AuthenticateUser: %v", err)
    }

    err = utils.CheckPasswordHash(string(user.Password), password)
    if err != nil {
        return nil, fmt.Errorf("AuthenticateUser: incorrect password")
    }

    return &user, nil
}


