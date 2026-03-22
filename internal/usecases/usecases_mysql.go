package usecases

import (
	"context"
	"rest-api/internal/models"
	"rest-api/userdb"

	"github.com/google/uuid"
)

const tagUsecaseMysql string = "UsecaseMsql"

type UsecasesMysql struct {
	service *userdb.ServiceUser
}

func NewMysql(serviceUser *userdb.ServiceUser) *UsecasesMysql {
	return &UsecasesMysql{
		service: serviceUser,
	}
}

func (u UsecasesMysql) ListUsers(ctx context.Context) ([]*models.User, error) {
	return u.service.ListUsers(ctx)
}

func (u UsecasesMysql) GetUserById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return u.service.GetUserById(ctx, id)
}

func (u UsecasesMysql) CreateUser(ctx context.Context, requestData models.CreateUserRequest) (*string, error) {
	return u.service.CreateUser(ctx, requestData)
}

func (u UsecasesMysql) UpdateUser(ctx context.Context, requestData models.UpdateUserRequest) (err error) {
	return u.service.UpdateUser(ctx, requestData)
}

func (u UsecasesMysql) DeleteUser(ctx context.Context, id uuid.UUID) (err error) {
	return u.service.DeleteUser(ctx, id)
}
