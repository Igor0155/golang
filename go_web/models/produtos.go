package models

import "github.com/Igor0155/golang/go_web/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	db := db.ConectaDB()

	selectAll, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}

	for selectAll.Next() {
		var id int
		var nome string
		var descricao string
		var preco float64
		var quantidade int
		err = selectAll.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaDB()

	insert, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = insert.Exec(nome, descricao, preco, quantidade)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
