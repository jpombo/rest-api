package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"rest-api/internal/models"
	"rest-api/internal/usecases"
	"rest-api/userdb"
)

const tagHandleMain string = "HandlerMain"

type Handlers struct {
	UsecaseUsers    *usecases.UsecasesUsers
	UsecaseProdutos *usecases.UsecasesProdutos
	ServiceUser     *userdb.ServiceUser
}

func New(usecaseUsers *usecases.UsecasesUsers, usecaseProdutos *usecases.UsecasesProdutos, serviceUser *userdb.ServiceUser) *Handlers {
	return &Handlers{
		UsecaseUsers:    usecaseUsers,
		UsecaseProdutos: usecaseProdutos,
		ServiceUser:     serviceUser,
	}
}

func (h Handlers) Listen(port int) error {

	if port <= 0 {
		slog.Error("invalid port", "port:", port)
		return errors.New("invalid port informed")
	} else {
		slog.Info("Listening into ", "port ", port)
	}

	h.registerUsersEndpoints()
	h.registerProdutosEndpoints()
	h.registerServiceUsersEndpoints()

	return http.ListenAndServe(
		fmt.Sprintf(":%v", port),
		nil,
	)
}

func (h Handlers) registerServiceUsersEndpoints() {
	http.HandleFunc("GET /usersdb", h.getAllUsersDB)
	http.HandleFunc("GET /useriddb", h.getUserIdDB)
	http.HandleFunc("POST /addusersdb", h.addUserDB)
	http.HandleFunc("POST /deluseriddb", h.deleteUserIdDB)
}

func (h Handlers) getAllUsersDB(w http.ResponseWriter, r *http.Request) {
	usersDB, err := h.ServiceUser.List(r.Context())
	w.WriteHeader(http.StatusOK)
	if err == nil {
		json.NewEncoder(w).Encode(usersDB)
	}
}
func (h Handlers) getUserIdDB(w http.ResponseWriter, r *http.Request) {
	var requestData models.GetUserResquest
	errRequestParse := json.NewDecoder(r.Body).Decode(&requestData)
	slog.Info(tagHandleMain, "requestParse", requestData)
	slog.Info(tagHandleMain, "errRequestParse", errRequestParse)
	if errRequestParse == nil {
		usersDB, err := h.ServiceUser.Get(r.Context(), requestData.UserID)
		slog.Info(tagHandleMain, "Response to usersIdDB", usersDB)
		w.WriteHeader(http.StatusOK)
		if err == nil {
			json.NewEncoder(w).Encode(usersDB)
		} else {
			slog.Error(tagHandleMain, "errDBResult", err)
		}
	} else {
		slog.Error(tagHandleMain, "errRequestParse", errRequestParse)
	}
}
func (h Handlers) addUserDB(w http.ResponseWriter, r *http.Request) {
	var requestData models.CreateUserRequest
	errRequestParse := json.NewDecoder(r.Body).Decode(&requestData)
	slog.Info(tagHandleMain, "requestParse", requestData)
	slog.Info(tagHandleMain, "errRequestParse", errRequestParse)
	if errRequestParse == nil {
		id, errAdd := h.ServiceUser.Create(r.Context(), requestData)
		slog.Info(tagHandleMain, "id", id, "errAdd", errAdd)
		if errAdd == nil {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(models.CreateUserResponse{
				NewUserID: id,
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
		slog.Info(tagHandleMain, "data received in wrong format", r.Body)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Reason: "invalid format",
		})
	}
}
func (h Handlers) deleteUserIdDB(w http.ResponseWriter, r *http.Request) {
	var requestData models.GetUserResquest
	errRequestParse := json.NewDecoder(r.Body).Decode(&requestData)
	slog.Info(tagHandleMain, "requestParse", requestData)
	slog.Info(tagHandleMain, "errRequestParse", errRequestParse)
	if errRequestParse == nil {
		_, err := h.ServiceUser.Get(r.Context(), requestData.UserID)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			slog.Info(tagHandleMain, "Response to deleteUsersIdDB", "user deleted")
		} else {
			w.WriteHeader(http.StatusNotModified)
			slog.Error(tagHandleMain, "errDBResult", err)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Reason: "Error",
			})
		}
	} else {
		slog.Error(tagHandleMain, "errRequestParse", errRequestParse)
	}
}
