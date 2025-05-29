package main

import (
	"github.com/Igor0155/golang/go_api/database"
	"github.com/Igor0155/golang/go_api/routes"
)

func main() {
	// models.Personalidades = []models.Personalidade{
	// 	{Id: 1, Nome: "Albert Einstein", Historia: "Físico teórico, desenvolveu a teoria da relatividade."},
	// 	{Id: 2, Nome: "Marie Curie", Historia: "Pioneira na pesquisa sobre radioatividade, primeira mulher a ganhar um Prêmio Nobel."},
	// 	{Id: 3, Nome: "Isaac Newton", Historia: "Matemático e físico, formulou as leis do movimento e da gravitação universal."},
	// }
	database.ConectarBanco()
	routes.HandleRequests()
}
