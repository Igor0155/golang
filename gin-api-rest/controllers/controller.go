package controllers

import (
	"net/http"

	"github.com/Igor0155/api-go-gin/database"
	"github.com/Igor0155/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func HelloWorldName( c *gin.Context){
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"Hello World, nome passado: " : nome,
	})

}

func ExibiTodosAluno(c *gin.Context){
	var alunos []models.Aluno

	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func ExibirAlunoPeloId(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Find(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Not Found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCpf(c *gin.Context){

	var aluno models.Aluno
	cpf := c.Params.ByName("cpf")

	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Not Found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": aluno})
}

func CriarNovoAluno(c *gin.Context){

	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func EditaAluno(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Find(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func DeletaAlunoPorId(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Find(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Not Found": "Aluno não encontrado"})
		return
	}
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno: "+ aluno.Nome +" deletado"})
}


