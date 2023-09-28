package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pos-graduacao-go-expert-desafio1/container"
	"github.com/pos-graduacao-go-expert-desafio1/internal/infrastructure/handlers"
	"github.com/pos-graduacao-go-expert-desafio1/internal/infrastructure/web"
)

func main() {
	// SQLITE3 CONNECTION
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// MYSQL CONNECTION
	// db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	handler := initHandlers(db)
	web.StartServer(handler)
}

func initHandlers(db *sql.DB) *handlers.CotationHandler {
	container := container.New(db)
	service := container.CotationService

	handler := handlers.NewCotationHandler(&service)

	return handler
}
