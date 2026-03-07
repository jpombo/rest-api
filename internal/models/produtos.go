package models

import "github.com/google/uuid"

// Entity Produto
type Produto struct {
	ID        uuid.UUID
	Descricao string
	Categoria string
}
type CreateProdutoRequest struct {
	Descricao string `json:"descricao"`
	Categoria string `json:"categoria"`
}
type CreateProdutoResponse struct {
	NewProdutoID uuid.UUID `json:"newProdutoId"`
}
