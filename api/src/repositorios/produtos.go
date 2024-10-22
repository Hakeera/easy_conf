package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
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

//Traz todos os produtos que atendem ao filtro
func (repositorio Produtos) Buscar(nomeProduto string) ([]modelos.Produto, error) {
	nomeProduto = fmt.Sprintf("%%%s%%", nomeProduto) // %nomeProduto% (dupla porcentagem faz pular a primeira porcentagem)

	linhas, erro := repositorio.db.Query(
		"select id, nome, custo, quantidade from produtos where nome LIKE ?", nomeProduto,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var produtos []modelos.Produto

	for linhas.Next() {
		var produto modelos.Produto

		if erro = linhas.Scan(
			&produto.ID,
			&produto.Nome,
			&produto.Custo,
			&produto.Quantidade,
		); erro != nil {
			return nil, erro
		}

		produtos = append(produtos, produto)
	}

	return produtos, nil
}