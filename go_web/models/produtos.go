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

	selectAll, err := db.Query("SELECT * FROM produtos order by id asc")
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
		p.Id = id
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

func DeletaProduto(id string) {
	db := db.ConectaDB()

	delete, err := db.Prepare("Delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}

	_, err = delete.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConectaDB()

	selectOne, err := db.Query("SELECT * FROM produtos WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	for selectOne.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectOne.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
	}

	defer db.Close()
	return p
}

func AtualizaProduto(id, nome, descricao string, preco float64, quantidade int) {

	db := db.ConectaDB()

	update, err := db.Prepare("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5")
	if err != nil {
		panic(err.Error())
	}

	_, err = update.Exec(nome, descricao, preco, quantidade, id)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
