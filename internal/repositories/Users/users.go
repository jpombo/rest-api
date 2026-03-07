package users

import (
	"log/slog"
	"rest-api/internal/models"
)

const tag string = "UserRepository"

type Users struct {
	users []models.User
}

func New() *Users {
	return &Users{
		users: make([]models.User, 0),
	}
}

func (u Users) GetAll() []models.User {
	return u.users
}
func (u Users) EmailExists(email string) (result bool) {
	result = false
	slog.Info(tag, "email", email, "userSize", len(u.users))
	for _, v := range u.users {
		slog.Info(tag, "emailMemory", v.Email)
		if v.Email == email {
			slog.Info(tag, "achada duplicidade", result)
			result = true
		}
	}
	return
}
func (u *Users) Add(newUser models.User) {
	u.users = append(u.users, newUser)
	slog.Info(tag, "Adicionado!", newUser)
}
