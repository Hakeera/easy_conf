package router

// O router.go cria um novo roteador e configura as rotas, enquanto o rotas.go define as rotas especificas para a manipulação de produtos, conctando-as às funções do controlador

// Ao iniciar o servidor HTTP no main.go o roteador configurado é usado, ele possui todas as rotas necessárias para manipular produtos (CRUD). Cada uma das rotas chama uma função no controlador correspondente, que contém a lógica para lidas com as operações

import (
	// Pacote que contém a definição das rotas
	"api/src/router/rotas"

	// Pacote para roteamento
	"github.com/gorilla/mux"
)

//Função que cria um novo roteador com o mux.NewRouter() e chama a função Configurar() para configurar e retornar o roteador
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}