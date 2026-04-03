package repositories

import (
	"rest-api/internal/models"
	produtos "rest-api/internal/repositories/Produtos"
	users "rest-api/internal/repositories/Users"
)

type Repositories struct {
	User interface {
		GetAll() []models.User
		Add(newUser models.User)
		EmailExists(email string) (result bool)
	}
	Produtos interface {
		ObterTodos() []models.Produto
		Adicionar(newProduto models.Produto)
		DescricaoExiste(descricao string) (result bool)
	}
}

func New() *Repositories {
	return &Repositories{
		User:     users.New(),
		Produtos: produtos.New(),
	}

}
