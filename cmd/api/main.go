package main

// Structure defined:
// handlers <-- usecases <-- repositories

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"rest-api/internal/handlers"
	"rest-api/internal/repositories"
	"rest-api/internal/usecases"
	"rest-api/userdb"
	"rest-api/userdb/db"

	_ "github.com/go-sql-driver/mysql"
)

const tagMain string = "Main"
const (
	dbUser         = "restapi"
	dbPassword     = "p@ssw0rD"
	database       = "restapi"
	dbRootPassword = "p@ssw0rD"
)

func main() {
	slog.Info("Starting API...")

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, "localhost", "3306", database)

	dbconn, err := sql.Open("mysql", dbUri)
	if err != nil {
		log.Fatal(err)
	}
	queries := db.New(dbconn)
	defer dbconn.Close()
	inicialize(queries)
}

func inicialize(queries *db.Queries) {
	repo := repositories.New()
	slog.Info(tagMain + " / Started repositories")

	useCaseUsers := usecases.NewUsers(repo)
	slog.Info(tagMain + " / Started usecase Users")
	useCaseProdutos := usecases.NewProdutos(repo)
	slog.Info(tagMain + " / Started usecase Produtos")
	useCaseMysql := usecases.NewMysql(userdb.NewServiceUser(queries))
	slog.Info(tagMain + " / Started usecase Mysql")

	h := handlers.New(useCaseUsers, useCaseProdutos, useCaseMysql)
	slog.Info(tagMain + " / Started handlers")

	h.Listen(8080)
}
