package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"strings"

	//"fmt"
	"io/ioutil"
	//"log"
	"net/http"
)

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

	if erro = produto.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

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
}

func BuscarProdutos(w http.ResponseWriter, r *http.Request) {
	nomeProduto := strings.ToLower(r.URL.Query().Get("produto"))

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