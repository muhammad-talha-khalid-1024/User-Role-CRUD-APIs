package domain

import (
	"errors"
	"time"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidUser       = errors.New("invalid user data")
)

type User struct {
	ID              string     `json:"id" db:"id"`
	Name            string     `json:"name" db:"name"`
	Email           string     `json:"email" db:"email"`
	EmailVerifiedAt *time.Time `json:"emailVerifiedAt" db:"emailVerifiedAt"`
	Password        string     `json:"-" db:"password"` // "-" hides from JSON
	Age             int        `json:"age" db:"age"`
	Status          int16      `json:"status" db:"status"`
	RememberToken   string     `json:"rememberToken" db:"rememberToken"`
	CreatedAt       time.Time  `json:"createdAt" db:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt" db:"updatedAt"`
	DeletedAt       time.Time  `json:"deletedAt" db:"deletedAt"`
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email,min=5,max=255"`
	Password string `json:"password" validate:"required,min=6,max=255"`
	Age      int    `json:"age" validate:"required,min=1,max=150"`
	Status   int16  `json:"status" validate:"required,min=1,max=2"`
}

type UpdateUserRequest struct {
	Name   string `json:"name" validate:"omitempty,min=2,max=100"`
	Email  string `json:"email" validate:"omitempty,email,min=5,max=255"`
	Age    int    `json:"age" validate:"omitempty,min=1,max=150"`
	Status int16  `json:"status" validate:"required,min=1,max=2"`
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id string) error
	List(limit, offset int) ([]*User, int, error)
}
