package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CriarProduto(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var produto modelos.Produto
	if erro = json.Unmarshal(corpoRequest, &produto); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioDeProdutos(db)
	produtoID, erro := repositorio.Criar(produto)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", produtoID)))
}

func BuscarProdutos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Produtos"))
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