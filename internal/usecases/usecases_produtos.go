package usecases

import (
	"errors"
	"log/slog"
	"rest-api/internal/models"
	"rest-api/internal/repositories"

	"github.com/google/uuid"
)

const tagUsecaseProdutos string = "UsecaseProdutos"
const errorDuplicateProduto string = "Produto already added"

type UsecasesProdutos struct {
	repos *repositories.Repositories
}

func NewProdutos(repos *repositories.Repositories) *UsecasesProdutos {
	return &UsecasesProdutos{
		repos: repos,
	}
}

func (up UsecasesProdutos) ObterTodos() []models.Produto {
	produtos := up.repos.Produtos.ObterTodos()
	return produtos
}
func (up UsecasesProdutos) Adicionar(newProduto models.CreateProdutoRequest) (id uuid.UUID, err error) {
	id = uuid.Nil
	err = nil

	slog.Info(tagUsecaseProdutos, "Validar", newProduto)
	if up.repos.Produtos.DescricaoExiste(newProduto.Descricao) == false {
		produtoToAdd := models.Produto{
			ID:        uuid.New(),
			Descricao: newProduto.Descricao,
			Categoria: newProduto.Categoria,
		}
		slog.Info(tagUsecaseProdutos, "Adicionar", newProduto)
		up.repos.Produtos.Adicionar(produtoToAdd)
		id = produtoToAdd.ID
	} else {
		slog.Info(tagUsecaseProdutos, "Descricao already added:", newProduto.Descricao)
		err = errors.New(errorDuplicateProduto)
	}
	return id, err
}
