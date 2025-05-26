package routes

import (
	"github.com/Igor0155/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET(":nome", controllers.HelloWorldName )
	r.GET("/alunos", controllers.ExibiTodosAluno)
	r.GET("/alunos/:id", controllers.ExibirAlunoPeloId)
	r.GET("alunos/cpf/:cpf", controllers.BuscaAlunoPorCpf)
	r.POST("/alunos", controllers.CriarNovoAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAlunoPorId)
	r.Run()
}