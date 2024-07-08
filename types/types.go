package types

import (
	"encoding/json"
	"time"
)

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=30"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
}

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ProductPayload struct {
	Name        string      `json:"name" validate:"required" `
	Description string      `json:"description" validate:"required"`
	Image       string      `json:"image" validate:"required"`
	Price       json.Number `json:"price" validate:"required,gt=0"`
	Quantity    json.Number `json:"quantity" validate:"required,gt=0"`
}

type ProductStore interface {
	GetProducts() ([]Product, error)
	CreateProduct(Product) error
}
