package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Representa todas as rotas da API
// URI: caminho da rota
// Metodo: O metodo HTTP (GET, POST, UPDATE, DELETE)
// Funcao: A função chamada quando a rota for acessada
type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
}

// Função que recebe um roteador (r) e configura as rotas
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasProdutos

	// Itera sobre um slice de rotas (rotasProdutos) e associa cada rota ao roteador usando o r.HandleFunc(), especificando a URI, a função de tratamento e o método HTTP
	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}