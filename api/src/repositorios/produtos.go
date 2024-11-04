package repositorios

// Responsavel pela lógica de acesso aos dados do banco de dados relacionados ao produtos

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Estrutura que representa o repositório de produtos. Contem uma referencia ao objeto de conexão com o banco de dados
type Produtos struct {
	db *sql.DB
}

// Cria um repositorio de produtos
func NovoRepositorioDeProdutos(db *sql.DB) *Produtos {
	return &Produtos{db}
}

// Insere um produto no banco de dados
func (repositorio Produtos) Criar(produto modelos.Produto) (uint64, error) {

	// Prepara o SQL para inserção do produto
	statement, erro := repositorio.db.Prepare(
		"insert into produtos (nome, custo, quantidade) values(?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	// Executa a declaração com os dados do produto
	resultado, erro := statement.Exec(produto.Nome, produto.Custo, produto.Quantidade)
	if erro != nil {
		return 0, erro
	}

	// Obtem o ID do último produto inserido
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	// Retorna o ID do produto ou um erro
	return uint64(ultimoIDInserido), nil
}

//Traz todos os produtos que atendem ao filtro
func (repositorio Produtos) Buscar(nomeProduto string) ([]modelos.Produto, error) {
	nomeProduto = fmt.Sprintf("%%%s%%", nomeProduto) // %nomeProduto% (dupla porcentagem faz pular a primeira porcentagem)

	// Formata o nome do produto para usar na consulta SQL com LIKE, que busca os produtos que tiverem o parametro passado
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

// Traz um usuario do banco de dados
func (repositorio Produtos) BuscarPorID(ID uint64) (modelos.Produto, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, custo, quantidade from produtos where id = ?",
		ID,
	)
	if erro != nil {
		return modelos.Produto{}, erro
	}
	defer linhas.Close()

	var produto modelos.Produto

	if linhas.Next() {
		if erro = linhas.Scan(
			&produto.ID,
			&produto.Nome,
			&produto.Custo,
			&produto.Quantidade,
		); erro != nil {
			return modelos.Produto{}, erro
		}
	}

	return produto, nil
}

func (repositorio Produtos) Atualizar(ID uint64, produto modelos.Produto) error {
	statement, erro := repositorio.db.Prepare("update produtos set nome = ?, custo = ?, quantidade = ? where id = ?",)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(produto.Nome, produto.Custo, produto.Quantidade, ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Produtos) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from produtos where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// O repositorio produtos possui a logica de interação com o banco de dados
// Fornece métodos para criar novos produtos e buscar produtos, mantendo a separação de preocupações entre a lógica de negócios (controladores) e a persistencia de dados (repositórios)

// O controlador CriarProduto chama o método Criar do repositório quando um novo produto é criado
// O controlador BuscarProdutos chama o método Buscar para recuperar produtos que correspondem ao nome fornecido na consulta