package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Responsavel por:
// Carregar as configurações essenciais a partir de variaveis de ambiente
// Configurar a porta do servidor e a string de conexão do banco de dados

// Quando o config.Carregar é chamada no main.go, todas essas configurações são definidas antes de iniciar o servidor, garantindo que a aplicação tenha as informações necessárias para funcionar corretamente

var (
	// Armazena a informação necessaria para conectar-se ao banco de dados
	StringConexaoBanco = ""
	// Um inteiro que armazenara o numero da porta em que a API ira escutar
	Porta              = 0
)

// Inicializar as variaveis de ambiente
func Carregar() {
	var erro error

	// godotenv.Load() tenta carregar as variaveis de ambiente de um arquivo .env. Se falhar, registra um erro e encerra a aplicação
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	// godotenv.Load("API_PORT") tenta obter a variavel de ambiente API_PORT
	// strconv.Atoi() converte essa string para um inteiro. Se não conseguir, define a porta padrão como 9000
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	// A StringConexaoBanco utiliza as variaveis de ambiente DB_USUARIO, DB_SENHA e DB_NOME para a aplicação conectar ao banco de dados
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"), 
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}