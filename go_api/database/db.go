package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectarBanco() {
	conect := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(conect))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados: " + err.Error())
	}
}
