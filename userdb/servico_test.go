package userdb_test

import (
	"context"
	"errors"
	"rest-api/internal/models"
	"rest-api/userdb"
	"rest-api/userdb/db"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-delve/delve/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)

const inputLayout = "2006-01-02"

func TestService_GetUser(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer d.Close()

	queries := db.New(d)
	id := uuid.New()
	dateConvert, err := time.Parse(inputLayout, "2026-01-01")
	t.Run("user encontrado", func(t *testing.T) {
		// fase: Arrange
		rows := sqlmock.NewRows([]string{"id", "name", "email", "birthdate"}).
			AddRow(id, "João de Teste", "jteste@test.com", dateConvert)
		mock.ExpectQuery("[A-Za-z]?select id, name, email, birthdate from tbusers where id = ?").
			WillReturnRows(rows)
		serviceTest := userdb.NewServiceUser(queries)
		// fase: Act
		found, _ := serviceTest.GetUserById(context.TODO(), id)

		// fase: Assert
		p := &models.User{
			ID:             id,
			Name:           "João de Teste",
			Email:          "jteste@test.com",
			DataNascimento: "2026-01-01",
		}
		assert.Nil(t, err)
		assert.Equal(t, p, found)
	})
	t.Run("user não encontrado", func(t *testing.T) {
		id := uuid.New()
		mock.ExpectQuery("[A-Za-z]?select id, name, email, birthdate from tbusers where id").WillReturnError(errors.New(""))
		service := userdb.NewServiceUser(queries)
		found, err := service.GetUserById(context.TODO(), id)
		assert.Nil(t, found)
		assert.Errorf(t, err, "erro lendo produto do repositório: %w")
	})
}
func TestService_GetProduct(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer d.Close()

	queries := db.New(d)
	id := uuid.New()
	t.Run("produto encontrado", func(t *testing.T) {
		// fase: Arrange
		rows := sqlmock.NewRows([]string{"id", "descricao", "categoria"}).
			AddRow(id, "Produto de Teste", "Categoria Teste")
		mock.ExpectQuery("[A-Za-z]?select id, descricao, categoria from tbproducts where id = ?").
			WillReturnRows(rows)
		serviceTest := userdb.NewServiceUser(queries)
		// fase: Act
		found, err := serviceTest.GetProductById(context.TODO(), id)

		// fase: Assert
		p := &models.Produto{
			ID:        id,
			Descricao: "Produto de Teste",
			Categoria: "Categoria Teste",
		}
		assert.Nil(t, err)
		assert.Equal(t, p, found)
	})
	t.Run("produto não encontrado", func(t *testing.T) {
		id := uuid.New()
		mock.ExpectQuery("[A-Za-z]?select id, descricao, categoria from tbproducts where id").WillReturnError(errors.New(""))
		service := userdb.NewServiceUser(queries)
		found, err := service.GetProductById(context.TODO(), id)
		assert.Nil(t, found)
		assert.Errorf(t, err, "erro lendo produto do repositório: %w")
	})
}
func TestCreateUserWithSQLMock(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer d.Close()
	name := "João de Teste"
	email := "jteste@test.com"
	datanascimento := "2026-01-01"

	mock.ExpectExec("[A-Za-z]?insert into tbusers").
		WithArgs(name, email, datanascimento).
		WillReturnResult(sqlmock.NewResult(1, 1))

	queries := db.New(d)
	service := userdb.NewServiceUser(queries)
	userTest := models.CreateUserRequest{
		Name:           name,
		Email:          email,
		DataNascimento: datanascimento,
	}
	id, err := service.CreateUser(context.TODO(), userTest)
	assert.NotNil(t, id)
}

func TestCreateProductWithSQLMock(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer d.Close()
	descricao := "Produto de Teste"
	categoria := "Categoria Teste"

	mock.ExpectExec("[A-Za-z]?insert into tbproducts").
		WithArgs(descricao, categoria).
		WillReturnResult(sqlmock.NewResult(1, 1))

	queries := db.New(d)
	service := userdb.NewServiceUser(queries)
	produtoTest := models.CreateProdutoRequest{
		Descricao: descricao,
		Categoria: categoria,
	}
	id, err := service.CreateProduct(context.TODO(), produtoTest)
	assert.NotNil(t, id)
}
