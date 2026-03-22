package userdb

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"rest-api/internal/models"
	"rest-api/userdb/db"
	"time"

	"github.com/google/uuid"
)

const tagServiceMysql string = "ServiceMsql"
const inputLayout = "2006-01-02"
const ErrorEmail = "E001"
const ErrorDate = "D001"
const ErrorDescricao = "DS01"

type ServiceUser struct {
	r *db.Queries
}

func NewServiceUser(r *db.Queries) *ServiceUser {
	return &ServiceUser{
		r: r,
	}
}

// Get a user
func (s *ServiceUser) GetUserById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	p, err := s.r.GetUser(ctx, id.String())
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	uuidValue, err := uuid.Parse(p.ID)
	if err == nil {
		return &models.User{
			ID:             uuidValue,
			Name:           p.Name,
			Email:          p.Email,
			DataNascimento: p.Birthdate.Format("2006-01-02"),
		}, nil
	} else {
		return nil, err
	}
}

// List user
func (s *ServiceUser) ListUsers(ctx context.Context) ([]*models.User, error) {
	result, err := s.r.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	slog.Info(tagServiceMysql, "result len", len(result))
	var users []*models.User
	for _, item := range result {
		uuidValue, err := uuid.Parse(item.ID)
		if err == nil {
			users = append(users, &models.User{
				ID:             uuidValue,
				Name:           item.Name,
				Email:          item.Email,
				DataNascimento: item.Birthdate.Format("2006-01-02"),
			})
		}
	}
	return users, nil
}

// Create a user
func (s *ServiceUser) CreateUser(ctx context.Context, user models.CreateUserRequest) (*string, error) {
	var id string
	var err error
	id = uuid.New().String()
	if s.EmailExist(ctx, user.Email) == false {
		slog.Info(tagServiceMysql, "uuid new value", id)
		dateValue, errParse := time.Parse(inputLayout, user.DataNascimento)
		if errParse != nil {
			err = errors.New(ErrorDate)
		} else {
			result, errCreate := s.r.CreateUser(ctx, db.CreateUserParams{
				ID:        id,
				Name:      user.Name,
				Email:     user.Email,
				Birthdate: dateValue,
			})
			if errCreate != nil {
				err = fmt.Errorf("error creating user: %w", err)
			} else if result != nil {
			}
		}
	} else {
		err = errors.New(ErrorEmail)
	}
	return &id, err
}

// Update user data
func (s *ServiceUser) UpdateUser(ctx context.Context, user models.UpdateUserRequest) (err error) {
	dateValue, errParse := time.Parse(inputLayout, user.DataNascimento)
	if errParse != nil {
		err = errors.New(ErrorDate)
	}
	rowsAffected, err := s.r.UpdateUser(ctx, db.UpdateUserParams{
		Name:      user.Name,
		Birthdate: dateValue,
		ID:        user.ID.String(),
	})
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	} else if rowsAffected == 0 {
		return fmt.Errorf("no row affected")
	}
	return nil
}

// Delete remove a user
func (s *ServiceUser) DeleteUser(ctx context.Context, id uuid.UUID) (err error) {
	rowsAffected, err := s.r.DeleteUser(ctx, id.String())
	if err != nil {
		return fmt.Errorf("error removing user: %w", err)
	} else if rowsAffected == 0 {
		return fmt.Errorf("no row affected")
	}
	return nil
}

func (s *ServiceUser) EmailExist(ctx context.Context, email string) (result bool) {
	result = false
	resultDB, _ := s.r.CheckEmail(ctx, email)
	if resultDB == 1 {
		result = true
	}
	return result
}

// Get a Product
func (s *ServiceUser) GetProductById(ctx context.Context, id uuid.UUID) (*models.Produto, error) {
	p, err := s.r.GetProduct(ctx, id.String())
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	uuidValue, err := uuid.Parse(p.ID)
	if err == nil {
		return &models.Produto{
			ID:        uuidValue,
			Descricao: p.Descricao,
			Categoria: p.Categoria,
		}, nil
	} else {
		return nil, err
	}
}

// List product
func (s *ServiceUser) ListProducts(ctx context.Context) ([]*models.Produto, error) {
	result, err := s.r.ListProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	slog.Info(tagServiceMysql, "result len", len(result))
	var products []*models.Produto
	for _, item := range result {
		uuidValue, err := uuid.Parse(item.ID)
		if err == nil {
			products = append(products, &models.Produto{
				ID:        uuidValue,
				Descricao: item.Descricao,
				Categoria: item.Categoria,
			})
		}
	}
	return products, nil
}

// Create a product
func (s *ServiceUser) CreateProduct(ctx context.Context, product models.CreateProdutoRequest) (*string, error) {
	var id string
	var err error
	id = uuid.New().String()
	if s.CheckDescricao(ctx, product.Descricao) == false {
		slog.Info(tagServiceMysql, "uuid new value", id)
		result, errCreate := s.r.CreateProduc(ctx, db.CreateProducParams{
			ID:        id,
			Descricao: product.Descricao,
			Categoria: product.Categoria,
		})
		if errCreate != nil {
			err = fmt.Errorf("error creating product: %w", err)
		} else if result != nil {
		}
	} else {
		err = errors.New(ErrorDescricao)
	}
	return &id, err
}

// Update product data
func (s *ServiceUser) UpdateProduct(ctx context.Context, product models.UpdateProductRequest) (err error) {
	rowsAffected, err := s.r.UpdateProduct(ctx, db.UpdateProductParams{
		Categoria: product.Descricao,
		ID:        product.ID.String(),
	})
	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	} else if rowsAffected == 0 {
		return fmt.Errorf("no row affected")
	}
	return nil
}

// Delete remove a user
func (s *ServiceUser) DeleteProduct(ctx context.Context, id uuid.UUID) (err error) {
	rowsAffected, err := s.r.DeleteProduct(ctx, id.String())
	if err != nil {
		return fmt.Errorf("error removing product: %w", err)
	} else if rowsAffected == 0 {
		return fmt.Errorf("no row affected")
	}
	return nil
}

func (s *ServiceUser) CheckDescricao(ctx context.Context, email string) (result bool) {
	result = false
	resultDB, _ := s.r.CheckDescricao(ctx, email)
	if resultDB == 1 {
		result = true
	}
	return result
}
