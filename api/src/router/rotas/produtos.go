package rotas

import (
	"api/src/controllers"
	"net/http"
)

// Um slice de Rota que define as rotas para os produtos
var rotasProdutos = []Rota{
	{
		URI:    "/produtos",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarProduto,
	},

	{
		URI:    "/produtos",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarProdutos,
	},

	{
		URI:    "/produtos/{produtoID}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarProduto,
	},

	{
		URI:    "/produtos/{produtoID}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarProduto,
	},

	{
		URI:    "/produtos/{produtoID}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarProduto,
	},
}