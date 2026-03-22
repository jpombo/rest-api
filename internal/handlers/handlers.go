package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"rest-api/internal/usecases"
)

const tagHandleMain string = "HandlerMain"

type Handlers struct {
	UsecaseUsers    *usecases.UsecasesUsers
	UsecaseProdutos *usecases.UsecasesProdutos
	UsecaseMysql    *usecases.UsecasesMysql
}

func New(usecaseUsers *usecases.UsecasesUsers, usecaseProdutos *usecases.UsecasesProdutos, usecaseMysql *usecases.UsecasesMysql) *Handlers {
	return &Handlers{
		UsecaseUsers:    usecaseUsers,
		UsecaseProdutos: usecaseProdutos,
		UsecaseMysql:    usecaseMysql,
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
	h.registerServiceProductsEndpoints()

	return http.ListenAndServe(
		fmt.Sprintf(":%v", port),
		nil,
	)
}
