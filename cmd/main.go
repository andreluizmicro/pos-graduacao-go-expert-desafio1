package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pos-graduacao-go-expert-desafio1/container"
	"github.com/pos-graduacao-go-expert-desafio1/internal/infrastructure/handlers"
	"github.com/pos-graduacao-go-expert-desafio1/internal/infrastructure/web"
)

func main() {
	// SQLITE3 CONNECTION
	db, err := sql.Open("sqlite3", "go-expert")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Para trabalhar com o banco de dados em memória criamos a tabela
	// Criar uma tabela em memória
	_, err = db.Exec(`CREATE TABLE cotation (
		id varchar(255),
		code varchar(80),
		code_in varchar(80),
		name varchar(255),
		high decimal(8,4),
		low decimal(8,4),
		var_bid decimal(8,4),
		pct_change decimal(8,4),
		bid decimal(8,4),
		ask decimal(8,4),
		created_at timestamp,
		primary key (id)
	)`)
	if err != nil {
		fmt.Println(err)
	}

	// Também podemos trabalhar com o banco MYSQL
	//MYSQL CONNECTION
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
