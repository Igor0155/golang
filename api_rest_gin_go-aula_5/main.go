package main

import (
	"github.com/Igor0155/golang/api_rest_gin_go-aula_5/database"
	"github.com/Igor0155/golang/api_rest_gin_go-aula_5/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
