package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DB interface {
	GetAllUsers() ([]User, error)
	GetUserByID(id string) (*User, error)
	CreateUser(user *User) error
}
