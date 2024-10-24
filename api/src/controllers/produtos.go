package controllers

// Contem a lógica para manipular os produtos
// Quando uma rota é acessada, o controlador correspondente é chamado

import (
	"api/src/banco" // Gerencia conexão com banco de dados
	"api/src/modelos" // Contém a definição dos modelos de dados(Produto)
	"api/src/repositorios" 
	"api/src/respostas" // Formata e envia respostas para o cliente
	"encoding/json" // Manipula dados JSON
	"strings" 

	//"fmt"
	"io/ioutil" // Lê o corpo da requisição
	//"log"
	"net/http" // Trabalha com requisições HTTP
)

// Lê o corpo da requisição 
func CriarProduto(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var produto modelos.Produto
	if erro = json.Unmarshal(corpoRequest, &produto); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Chama Preparar() no modelo Produto para validas ou preparar os dados
	if erro = produto.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Conecta ao banco de dados e cria um novo produto através de repositório
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeProdutos(db)
	produto.ID, erro = repositorio.Criar(produto)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, produto)

	// Se tudo ocorrer bem retorna 201, caso contrário retorna o erro
}

// Obtem o nome do produto a partir da query string
func BuscarProdutos(w http.ResponseWriter, r *http.Request) {
	nomeProduto := strings.ToLower(r.URL.Query().Get("produto"))

	// Conecta ao banco de dados e busca produtos utilizando o repositório
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeProdutos(db)
	produtos, erro := repositorio.Buscar(nomeProduto)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, produtos)

	// Retorna os produtos encontrado e um 200, se ocorrer erro retorna o erro
}

func BuscarProduto(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Produto"))
}

func AtualizarProduto(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Produto"))
}

func DeletarProduto(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Produto"))
}