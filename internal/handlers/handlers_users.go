package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"rest-api/internal/models"
)

const tagHandleUser string = "HandlerUser"

func (h Handlers) registerUsersEndpoints() {
	http.HandleFunc("GET /users", h.getAllUsers)
	http.HandleFunc("POST /addusers", h.addUser)
}

func (h Handlers) getAllUsers(w http.ResponseWriter, r *http.Request) {
	slog.Info(tagHandleUser, "Response to GetAllUser", h.UsecaseUsers.GetAll())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.UsecaseUsers.GetAll())
}

func (h Handlers) addUser(w http.ResponseWriter, r *http.Request) {
	var requestData models.CreateUserRequest
	errRequestParse := json.NewDecoder(r.Body).Decode(&requestData)
	slog.Info(tagHandleUser, "requestParse", requestData)
	slog.Info(tagHandleUser, "errRequestParse", errRequestParse)
	if errRequestParse == nil {
		id, errAdd := h.UsecaseUsers.Add(requestData)
		slog.Info(tagHandleUser, "id", id, "errAdd", errAdd)
		if errAdd == nil {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.CreateUserResponse{
				NewUserID: id,
			})
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Reason: "email duplicated",
			})
		}
	} else {
		slog.Info(tagHandleUser, "data received in wrong format", r.Body)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Reason: "invalid format",
		})
	}
}
