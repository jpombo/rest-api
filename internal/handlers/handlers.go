package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"rest-api/internal/usecases"
)

type Handlers struct {
	UsecaseUsers    *usecases.UsecasesUsers
	UsecaseProdutos *usecases.UsecasesProdutos
}

func New(usecaseUsers *usecases.UsecasesUsers, usecaseProdutos *usecases.UsecasesProdutos) *Handlers {
	return &Handlers{
		UsecaseUsers:    usecaseUsers,
		UsecaseProdutos: usecaseProdutos,
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

	return http.ListenAndServe(
		fmt.Sprintf(":%v", port),
		nil,
	)
}
