package models

import (
	"time"

	"github.com/google/uuid"
)

// struct for auth.register request
type AuthRegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required,min=2"`
}

type AuthRegisterUser struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// struct for auth.register response
type AuthRegisterResponse struct {
	Message string           `json:"message"`
	Status  string           `json:"status"`
	Code    int              `json:"code"`
	User    AuthRegisterUser `json:"user"`
}
