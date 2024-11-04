package respostas

// Formata e envia respostas JSON para o cliente

import (
	"encoding/json"
	"log"
	"net/http"
)

//Retorna uma resposta em json para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

//Retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}

// Centraliza a lógica para enviar respostas ao cliente, tanto para sucesso quanto para erro
// É usado em todo o controlador para enviar respostas apropriadas ao cliente após as operações, como criar ou buscar produtos
// Se ocorrer um erro ao criar um produto a função CriarProduto chama respostas.Erro() para enviar a resposta apropriada ao cliente