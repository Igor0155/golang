package main

import (
	"github.com/Igor0155/api-go-gin/database"
	"github.com/Igor0155/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()

	// models.Alunos = []models.Aluno{
	// 	{Nome: "Igor", CPF: "12345678912", RG: "114516843"},
	// 	{Nome: "Victor", CPF: "78945645612", RG: "135138381"},
	// 	{Nome: "Rod", CPF: "45612348912", RG: "1351384383"},
	// 	{Nome: "Deb", CPF: "12345678945", RG: "513813813"},
	// 	{Nome: "Pedro", CPF: "42613218749", RG: "78974564123"},
	// }
}