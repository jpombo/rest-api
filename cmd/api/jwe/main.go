package main

// Structure defined:
// handlers <-- usecases <-- repositories

import (
	"log"
	"log/slog"
	"net/http"
	"rest-api/internal/handlers"
	"rest-api/internal/repositories"
	"rest-api/internal/usecases"

	_ "github.com/go-sql-driver/mysql"
)

const tagMain string = "MainJWE"

func main() {
	slog.Info("Starting API...")
	startJWE()
}

func startJWE() {
	// key, err := repositories.GenerateKey()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	key := []byte("12345678901234567890123456789012")
	repoJWE := repositories.NewCryptoRepository(key)
	usecase := usecases.NewCryptoUseCase(repoJWE)
	handler := handlers.NewCryptoHandler(usecase)

	http.HandleFunc("/encrypt", handler.Encrypt)
	http.HandleFunc("/decrypt", handler.Decrypt)

	log.Println("Servidor JWE rodando em :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
