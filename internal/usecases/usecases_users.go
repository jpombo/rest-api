package usecases

import (
	"errors"
	"log/slog"
	"rest-api/internal/models"
	"rest-api/internal/repositories"

	"github.com/google/uuid"
)

const tagUsecaseUsers string = "UsecaseUsers"
const errorDuplicateUser string = "Email already added"

type UsecasesUsers struct {
	repos *repositories.Repositories
}

func NewUsers(repos *repositories.Repositories) *UsecasesUsers {
	return &UsecasesUsers{
		repos: repos,
	}
}

func (u UsecasesUsers) GetAll() []models.User {
	users := u.repos.User.GetAll()
	return users
}
func (u UsecasesUsers) Add(newUser models.CreateUserRequest) (id uuid.UUID, err error) {
	id = uuid.Nil
	err = nil

	slog.Info(tagUsecaseUsers, "Add", newUser)
	if u.repos.User.EmailExists(newUser.Email) == false {
		userToAdd := models.User{
			ID:    uuid.New(),
			Name:  newUser.Name,
			Email: newUser.Email,
		}
		u.repos.User.Add(userToAdd)
		id = userToAdd.ID
	} else {
		slog.Info(tagUsecaseUsers, "Email already added:", newUser.Email)
		err = errors.New(errorDuplicateUser)
	}
	return id, err
}
