package main

// Structure defined:
// handlers <-- usecases <-- repositories

import (
	"log/slog"
	"rest-api/internal/handlers"
	"rest-api/internal/repositories"
	"rest-api/internal/usecases"
)

const tagMain string = "Main"

func main() {
	slog.Info("Starting API...")
	inicialize()
}

func inicialize() {
	repo := repositories.New()
	slog.Info(tagMain + " / Started repositories")

	useCaseUsers := usecases.NewUsers(repo)
	slog.Info(tagMain + " / Started usecase Users")
	useCaseProdutos := usecases.NewProdutos(repo)
	slog.Info(tagMain + " / Started usecase Produtos")

	h := handlers.New(useCaseUsers, useCaseProdutos)
	slog.Info(tagMain + " / Started handlers")

	h.Listen(8080)
}

func inicialize1() {
	repo := repositories.New()
	handlers.New(usecases.NewUsers(repo), usecases.NewProdutos(repo)).Listen(8080)
}
