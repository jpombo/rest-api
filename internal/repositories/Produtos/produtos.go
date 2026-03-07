package produtos

import (
	"log/slog"
	"rest-api/internal/models"
)

const tag string = "ProdutoRepository"

type Produtos struct {
	produtos []models.Produto
}

func New() *Produtos {
	return &Produtos{
		produtos: make([]models.Produto, 0),
	}
}

func (p Produtos) ObterTodos() []models.Produto {
	return p.produtos
}

func (p Produtos) DescricaoExiste(descricao string) (result bool) {
	result = false
	slog.Info(tag, "descricao", descricao, "ProdutoSize", len(p.produtos))
	for _, v := range p.produtos {
		slog.Info(tag, "descricaoMemory", v.Descricao)
		if v.Descricao == descricao {
			slog.Info(tag, "achada duplicidade", result)
			result = true
		}
	}
	return
}

func (p *Produtos) Adicionar(newProduto models.Produto) {
	p.produtos = append(p.produtos, newProduto)
	slog.Info(tag, "Adicionado!", newProduto)
}
