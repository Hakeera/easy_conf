package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Repositorio de produtos
type Produtos struct {
	db *sql.DB
}

// Cria um repositorio de produtos
func NovoRepositorioDeProdutos(db *sql.DB) *Produtos {
	return &Produtos{db}
}

// Insere um produto no banco de dados
func (repositorio Produtos) Criar(produto modelos.Produto) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into produtos (nome, custo, quantidade) values(?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(produto.Nome, produto.Custo, produto.Quantidade)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}