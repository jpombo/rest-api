package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"rest-api/internal/models"
)

const tagHandleMysql string = "HandlerMysql"

func (h Handlers) registerServiceUsersEndpoints() {
	http.HandleFunc("GET /usersdb", h.getAllUsersDB)
	http.HandleFunc("GET /useriddb", h.getUserIdDB)
	http.HandleFunc("POST /addusersdb", h.addUserDB)
	http.HandleFunc("POST /updateusersdb", h.updateUserDB)
	http.HandleFunc("POST /deluseriddb", h.deleteUserIdDB)
}
func (h Handlers) getAllUsersDB(w http.ResponseWriter, r *http.Request) {
	usersDB, err := h.UsecaseMysql.ListUsers(r.Context())
	w.WriteHeader(http.StatusOK)
	if err == nil {
		json.NewEncoder(w).Encode(usersDB)
	}
}
func (h Handlers) getUserIdDB(w http.ResponseWriter, r *http.Request) {
	var requestData models.GetUserResquest
	errRequestParse := json.NewDecoder(r.Body).Decode(&requestData)
	slog.Info(tagHandleMysql, "requestParse", requestData)
	slog.Info(tagHandleMysql, "errRequestParse", errRequestParse)
	if errRequestParse == nil {
		usersDB, err := h.UsecaseMysql.GetUserById(r.Context(), requestData.UserID)
		slog.Info(tagHandleMysql, "Response to usersIdDB", usersDB)
		w.WriteHeader(http.StatusOK)
		if err == nil {
			json.NewEncoder(w).Encode(usersDB)
		} else {
			slog.Error(tagHandleMysql, "errDBResult", err)
		}
	} else {
		slog.Error(tagHandleMysql, "errRequestParse", errRequestParse)
	}
}
func (h Handlers) addUserDB(w http.ResponseWriter, r *http.Request) {
	var requestData models.CreateUserRequest
	errRequestParse := json.NewDecoder(r.Body).Decode(&requestData)
	slog.Info(tagHandleMysql, "requestParse", requestData)
	slog.Info(tagHandleMysql, "errRequestParse", errRequestParse)
	if errRequestParse == nil {
		id, errAdd := h.UsecaseMysql.CreateUser(r.Context(), requestData)
		slog.Info(tagHandleMysql, "id", *id, "errAdd", errAdd)
		if errAdd == nil {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.CreateUserResponse{
				NewUserID: *id,
			})
		} else {
			errorText := fmt.Sprint("add User error code:", errAdd)
			switch errAdd.Error() {
			case "E001":
				w.WriteHeader(http.StatusConflict)
			case "D001":
				w.WriteHeader(http.StatusBadRequest)
			}
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Reason: errorText,
			})
		}
	} else {
		slog.Info(tagHandleMysql, "data received in wrong format", r.Body)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Reason: "invalid format",
		})
	}
}
func (h Handlers) updateUserDB(w http.ResponseWriter, r *http.Request) {
	var requestData models.UpdateUserRequest
	errRequestParse := json.NewDecoder(r.Body).Decode(&requestData)
	slog.Info(tagHandleMysql, "requestParse", requestData)
	slog.Info(tagHandleMysql, "errRequestParse", errRequestParse)
	if errRequestParse == nil {
		err := h.UsecaseMysql.UpdateUser(r.Context(), requestData)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			slog.Info(tagHandleMysql, "Response to updateUsersIdDB", "user updated")
		} else if err.Error() == "no row affected" {
			w.WriteHeader(http.StatusNoContent)
			slog.Info(tagHandleMysql, "DBResult", err)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Reason: "No affected",
			})
		} else {
			w.WriteHeader(http.StatusNotModified)
			slog.Error(tagHandleMysql, "errDBResult", err)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Reason: "Error",
			})
		}
	} else {
		slog.Error(tagHandleMysql, "errRequestParse", errRequestParse)
	}
}
func (h Handlers) deleteUserIdDB(w http.ResponseWriter, r *http.Request) {
	var requestData models.GetUserResquest
	errRequestParse := json.NewDecoder(r.Body).Decode(&requestData)
	slog.Info(tagHandleMysql, "requestParse", requestData)
	slog.Info(tagHandleMysql, "errRequestParse", errRequestParse)
	if errRequestParse == nil {
		err := h.UsecaseMysql.DeleteUser(r.Context(), requestData.UserID)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			slog.Info(tagHandleMysql, "Response to deleteUsersIdDB", "user deleted")
		} else if err.Error() == "no row affected" {
			w.WriteHeader(http.StatusNoContent)
			slog.Info(tagHandleMysql, "DBResult", err)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Reason: "No affected",
			})
		} else {
			w.WriteHeader(http.StatusNotModified)
			slog.Error(tagHandleMysql, "errDBResult", err)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Reason: "Error",
			})
		}
	} else {
		slog.Error(tagHandleMysql, "errRequestParse", errRequestParse)
	}
}

func (h Handlers) registerServiceProductsEndpoints() {
	http.HandleFunc("GET /productsdb", h.getAllProductsDB)
	http.HandleFunc("GET /productiddb", h.getProductIdDB)
	http.HandleFunc("POST /addproductsdb", h.addProductDB)
	http.HandleFunc("POST /updateproductsdb", h.updateProductDB)
	http.HandleFunc("POST /delproductiddb", h.deleteProductIdDB)
}

func (h Handlers) getAllProductsDB(w http.ResponseWriter, r *http.Request) {
	productsDB, err := h.UsecaseMysql.ListProduct(r.Context())
	w.WriteHeader(http.StatusOK)
	if err == nil {
		json.NewEncoder(w).Encode(productsDB)
	}
}
func (h Handlers) getProductIdDB(w http.ResponseWriter, r *http.Request) {
	var requestData models.GetProductResquest
	errRequestParse := json.NewDecoder(r.Body).Decode(&requestData)
	slog.Info(tagHandleMysql, "requestParse", requestData)
	slog.Info(tagHandleMysql, "errRequestParse", errRequestParse)
	if errRequestParse == nil {
		productsDB, err := h.UsecaseMysql.GetProductById(r.Context(), requestData.ProductID)
		slog.Info(tagHandleMysql, "Response to productsIdDB", productsDB)
		w.WriteHeader(http.StatusOK)
		if err == nil {
			json.NewEncoder(w).Encode(productsDB)
		} else {
			slog.Error(tagHandleMysql, "errDBResult", err)
		}
	} else {
		slog.Error(tagHandleMysql, "errRequestParse", errRequestParse)
	}
}
func (h Handlers) addProductDB(w http.ResponseWriter, r *http.Request) {
	var requestData models.CreateProdutoRequest
	errRequestParse := json.NewDecoder(r.Body).Decode(&requestData)
	slog.Info(tagHandleMysql, "requestParse", requestData)
	slog.Info(tagHandleMysql, "errRequestParse", errRequestParse)
	if errRequestParse == nil {
		id, errAdd := h.UsecaseMysql.CreateProduct(r.Context(), requestData)
		slog.Info(tagHandleMysql, "id", *id, "errAdd", errAdd)
		if errAdd == nil {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.CreateProdutoResponse{
				NewProdutoID: *id,
			})
		} else {
			errorText := fmt.Sprint("add Product error code:", errAdd)
			if errAdd.Error() == "DS01" {
				w.WriteHeader(http.StatusConflict)
			}
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Reason: errorText,
			})
		}
	} else {
		slog.Info(tagHandleMysql, "data received in wrong format", r.Body)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Reason: "invalid format",
		})
	}
}
func (h Handlers) updateProductDB(w http.ResponseWriter, r *http.Request) {
	var requestData models.UpdateProductRequest
	errRequestParse := json.NewDecoder(r.Body).Decode(&requestData)
	slog.Info(tagHandleMysql, "requestParse", requestData)
	slog.Info(tagHandleMysql, "errRequestParse", errRequestParse)
	if errRequestParse == nil {
		err := h.UsecaseMysql.UpdateProduct(r.Context(), requestData)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			slog.Info(tagHandleMysql, "Response to updateProductIdDB", "product updated")
		} else if err.Error() == "no row affected" {
			w.WriteHeader(http.StatusNoContent)
			slog.Info(tagHandleMysql, "DBResult", err)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Reason: "No affected",
			})
		} else {
			w.WriteHeader(http.StatusNotModified)
			slog.Error(tagHandleMysql, "errDBResult", err)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Reason: "Error",
			})
		}
	} else {
		slog.Error(tagHandleMysql, "errRequestParse", errRequestParse)
	}
}
func (h Handlers) deleteProductIdDB(w http.ResponseWriter, r *http.Request) {
	var requestData models.GetProductResquest
	errRequestParse := json.NewDecoder(r.Body).Decode(&requestData)
	slog.Info(tagHandleMysql, "requestParse", requestData)
	slog.Info(tagHandleMysql, "errRequestParse", errRequestParse)
	if errRequestParse == nil {
		err := h.UsecaseMysql.DeleteProduct(r.Context(), requestData.ProductID)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			slog.Info(tagHandleMysql, "Response to deleteProductIdDB", "products deleted")
		} else if err.Error() == "no row affected" {
			w.WriteHeader(http.StatusNoContent)
			slog.Info(tagHandleMysql, "DBResult", err)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Reason: "No affected",
			})
		} else {
			w.WriteHeader(http.StatusNotModified)
			slog.Error(tagHandleMysql, "errDBResult", err)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Reason: "Error",
			})
		}
	} else {
		slog.Error(tagHandleMysql, "errRequestParse", errRequestParse)
	}
}
