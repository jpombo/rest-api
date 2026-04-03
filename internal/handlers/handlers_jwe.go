package handlers

import (
	"encoding/json"
	"net/http"
	"rest-api/internal/models"
	"rest-api/internal/usecases"
)

const tagHandleJwe string = "HandlerJwe"

type CryptoHandler struct {
	usecase usecases.CryptoUseCase
}

func NewCryptoHandler(u usecases.CryptoUseCase) *CryptoHandler {
	return &CryptoHandler{usecase: u}
}

func (h *CryptoHandler) Encrypt(w http.ResponseWriter, r *http.Request) {
	var req models.EncryptRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.usecase.EncryptJSON(req.Payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(models.EncryptResponse{JWE: result})
}

func (h *CryptoHandler) Decrypt(w http.ResponseWriter, r *http.Request) {
	var req models.DecryptRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.usecase.DecryptJWE(req.JWE)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(models.DecryptResponse{Payload: result})
}
