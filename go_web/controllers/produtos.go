package controllers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/Igor0155/golang/go_web/models"
)

var templateHTML = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosProdutos()
	templateHTML.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {

	templateHTML.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		panic("ID não informado")
	}
	produto := models.EditaProduto(id)

	templateHTML.ExecuteTemplate(w, "Edit", produto)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	quantidade := r.FormValue("quantidade")

	precoConvertido, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		panic(err.Error())
	}

	quantidadeConvertido, err := strconv.Atoi(quantidade)
	if err != nil {
		panic(err.Error())
	}

	models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertido)

	http.Redirect(w, r, "/", 301)
	return
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	id := r.FormValue("id")
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	quantidade := r.FormValue("quantidade")

	precoConvertido, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		panic(err.Error())
	}

	quantidadeConvertido, err := strconv.Atoi(quantidade)
	if err != nil {
		panic(err.Error())
	}

	models.AtualizaProduto(id, nome, descricao, precoConvertido, quantidadeConvertido)

	http.Redirect(w, r, "/", 301)
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		panic("ID não informado")
	}
	models.DeletaProduto(id)
	http.Redirect(w, r, "/", 301)

	return
}
