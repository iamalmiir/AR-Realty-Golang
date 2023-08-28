package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateUser(user *User) error {
	sqlStatement := `INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(sqlStatement, user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)

	return err
}

func GetUser(id uuid.UUID) (User, error) {
	var user User
	sqlStatement := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = ? LIMIT 1;`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, nil // Return empty user if the user doesn't exist
		}
		return User{}, err
	}

	return user, nil
}

func CheckEmail(email string, user *User) bool {
	sqlStatement := `SELECT id, name, email, password FROM users WHERE email = ? LIMIT 1`

	err := db.QueryRow(sqlStatement, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return false // No matching user found
		}
		return false // Other error occurred
	}

	return true
}
