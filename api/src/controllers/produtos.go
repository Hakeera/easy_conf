package controllers

import "net/http"

func CriarProduto(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Produto"))
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