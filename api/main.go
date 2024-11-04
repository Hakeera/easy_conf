package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// FLUXO DO PROGRAMA
// 1. As configurações são carregadas
// 2. Um roteador é gerado para definir as rotas
// 3. O servidor HTTP é iniciado e fica escutando em uma porta especificada
// 4. Caso haja algum erro ao levantar o servidor, o programa é iniciado e o erro é registrado no log.

func main() {
	//Carrega configurações da aplicação, como a porta do servidor ou credenciais de banco de dados. Possivelmente informações de um arquivo de configuração ou de variaveis de ambiente
	config.Carregar()

	// Cria e retorna um router, que define como cada rota HTTP sera tratada
	r := router.Gerar()

	// Levanta o servidor HTTP na porta especificada e associa o router retornado pela função router.Gerar()
	fmt.Printf("Escutando na porta %d", config.Porta)
	// Se ocorrer um erro, o programa sera finalizado e o erro sera exibido
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}