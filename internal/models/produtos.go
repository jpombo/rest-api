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
	NewProdutoID string `json:"newProdutoId"`
}
type GetProductResquest struct {
	ProductID uuid.UUID `json:"productId"`
}
type UpdateProductRequest struct {
	ID        uuid.UUID `json:"productId"`
	Descricao string    `json:"categoria"`
}
