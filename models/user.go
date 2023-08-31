package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserStorage struct {
	Conn *sqlx.DB
}

type User struct {
	ID        uuid.UUID `db:"id" json:"id"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// Create new user
func (s *UserStorage) NewUser(user *User) error {
	query := `
		INSERT INTO users (id, first_name, last_name, email, password, created_at, updated_at) 
		VALUES (:id, :first_name, :last_name, :email, :password, :created_at, :updated_at)
	`

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := s.Conn.NamedExec(query, user)
	if err != nil {
		return err
	}

	return nil
}

// Get all users
func (s *UserStorage) GetUsers() ([]User, error) {
	query := `
        SELECT id, first_name, last_name, email, password, created_at, updated_at
        FROM users
    `

	var users []User
	err := s.Conn.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}
