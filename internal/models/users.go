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
	NewUserID uuid.UUID `json:"newUserId"`
}
type GetUserResquest struct {
	UserID uuid.UUID `json:"userId"`
}
