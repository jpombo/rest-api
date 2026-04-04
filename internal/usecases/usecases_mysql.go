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

// User
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

// Product
func (u UsecasesMysql) ListProduct(ctx context.Context) ([]*models.Produto, error) {
	return u.service.ListProducts(ctx)
}
func (u UsecasesMysql) GetProductById(ctx context.Context, id uuid.UUID) (*models.Produto, error) {
	return u.service.GetProductById(ctx, id)
}
func (u UsecasesMysql) CreateProduct(ctx context.Context, requestData models.CreateProdutoRequest) (*string, error) {
	return u.service.CreateProduct(ctx, requestData)
}
func (u UsecasesMysql) UpdateProduct(ctx context.Context, requestData models.UpdateProductRequest) (err error) {
	return u.service.UpdateProduct(ctx, requestData)
}
func (u UsecasesMysql) DeleteProduct(ctx context.Context, id uuid.UUID) (err error) {
	return u.service.DeleteProduct(ctx, id)
}
