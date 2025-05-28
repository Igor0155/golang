package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaDB() *sql.DB {
	conexao := "user=postgres password=admin dbname=alura_loja port=5432 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
