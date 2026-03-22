package models

import "github.com/google/uuid"

// Entity User
type User struct {
	ID             uuid.UUID
	Name           string
	Email          string
	DataNascimento string
}
type CreateUserRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	DataNascimento string `json:"datanasc"`
}
type CreateUserResponse struct {
	NewUserID string `json:"newUserId"`
}
type GetUserResquest struct {
	UserID uuid.UUID `json:"userId"`
}
type UpdateUserRequest struct {
	ID             uuid.UUID `json:"userId"`
	Name           string    `json:"name"`
	DataNascimento string    `json:"datanasc"`
}
