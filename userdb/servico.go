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

const tagServiceUser string = "ServiceUser"
const inputLayout = "2006-01-02"
const ErrorEmail = "E001"
const ErrorDate = "D001"

type ServiceUser struct {
	r *db.Queries
}

func NewServiceUser(r *db.Queries) *ServiceUser {
	return &ServiceUser{
		r: r,
	}
}

// Get a user
func (s *ServiceUser) Get(ctx context.Context, id uuid.UUID) (*models.User, error) {
	p, err := s.r.Get(ctx, id.String())
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	uuidValue, err := uuid.Parse(p.ID)
	if err == nil {
		return &models.User{
			ID:             uuidValue,
			Name:           p.Name,
			Email:          p.Email,
			DataNascimento: p.Birthdate.Format("dd/mm/yyyy"),
		}, nil
	} else {
		return nil, err
	}
}

// List user
func (s *ServiceUser) List(ctx context.Context) ([]*models.User, error) {
	result, err := s.r.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error reading from database: %w", err)
	}
	slog.Info(tagServiceUser, "result len", len(result))
	var users []*models.User
	for _, item := range result {
		uuidValue, err := uuid.Parse(item.ID)
		if err == nil {
			users = append(users, &models.User{
				ID:             uuidValue,
				Name:           item.Name,
				Email:          item.Email,
				DataNascimento: item.Birthdate.Format(inputLayout),
			})
		}
	}
	return users, nil
}

// Create a user
func (s *ServiceUser) Create(ctx context.Context, user models.CreateUserRequest) (id uuid.UUID, err error) {
	id = uuid.Nil
	err = nil
	if s.EmailExist(ctx, user.Email) == false {
		id := uuid.New()
		slog.Info(tagServiceUser, "uuid new value", id)
		dateValue, errParse := time.Parse(inputLayout, user.DataNascimento)
		if errParse != nil {
			err = errors.New(ErrorDate)
		} else {
			result, errCreate := s.r.Create(ctx, db.CreateParams{
				ID:        id.String(),
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
	return id, err
}

// // Update user data
// func (s *ServiceUser) Update(ctx context.Context, user *models.User) error {
// 	uuidValue := []byte(user.ID.String())
// 	dateValue, err := time.Parse("dd/mm/yyyy", user.DataNascimento)
// 	err := s.r.Update(ctx, db.UpdateParams{
// 		Name:      user.Name,
// 		Email:     user.Email,
// 		Birthdate: dateValue,
// 		ID:        uuidValue,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("error updating person: %w", err)
// 	}
// 	return nil
// }

// Delete remove a user
func (s *ServiceUser) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.r.Delete(ctx, id.String())
	if err != nil {
		return fmt.Errorf("error removing user: %w", err)
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
