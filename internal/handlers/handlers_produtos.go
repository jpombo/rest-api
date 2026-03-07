package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"rest-api/internal/models"
)

const tagHandleProdutos string = "HandlerProdutos"

func (h Handlers) registerProdutosEndpoints() {
	http.HandleFunc("GET /produtos", h.obterTodosProdutos)
	http.HandleFunc("POST /adicionaprodutos", h.adicionarProdutos)
}

func (h Handlers) obterTodosProdutos(w http.ResponseWriter, r *http.Request) {
	slog.Info(tagHandleUser, "Response:", h.UsecaseProdutos.ObterTodos())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.UsecaseProdutos.ObterTodos())
}

func (h Handlers) adicionarProdutos(w http.ResponseWriter, r *http.Request) {
	var produtoRequest models.CreateProdutoRequest
	errRequestParse := json.NewDecoder(r.Body).Decode(&produtoRequest)
	slog.Info(tagHandleProdutos, "requestParse", produtoRequest)
	slog.Info(tagHandleProdutos, "errRequestParse", errRequestParse)
	if errRequestParse == nil {
		id, errAdd := h.UsecaseProdutos.Adicionar(produtoRequest)
		slog.Info(tagHandleProdutos, "id", id, "errAdd", errAdd)
		if errAdd == nil {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.CreateProdutoResponse{
				NewProdutoID: id,
			})
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Reason: "descricao duplicated",
			})
		}
	} else {
		slog.Info(tagHandleProdutos, "data received in wrong format", r.Body)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Reason: "invalid format",
		})
	}
}
