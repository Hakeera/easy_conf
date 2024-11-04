package modelos

// Define a estrutura do modelo de dados Produto e inclui metodos para validas e formatar os dados do produto

import (
	"errors"
	"strings"
)

//Atributos dos produtos
type Produto struct {
	// Omitempty faz com que o campo não seja incluso no JSON caso seu valor seja vazio
	ID uint64 `json:"id,omitempty"` 
	Nome string `json:"nome,omitempty"`
	Custo float64 `json:"custo,omitempty"`
	Quantidade int `json:"quantidade,omitempty"`
}


// Verifica se os campos obrigatórios estão preenchidos
func (produto *Produto) validar() error {
	if produto.Nome == "" {
		return errors.New("o nome do produto é obrigatório")
	}

	if produto.Custo <= 0 {
		return errors.New("o custo do produto é obrigatório")
	}

	if produto.Quantidade <= 0 {
		return errors.New("a quantidade do produto é obrigatória")
	}

	return nil
}

// Remove espaços em branco
func (produto *Produto) formatar() {
	produto.Nome = strings.TrimSpace(produto.Nome)
}

//Chama os metodos para validar e formatar o produto
func (produto *Produto) Preparar() error {
	if erro := produto.validar(); erro != nil {
		return erro
	}

	produto.formatar()
	return nil
}

// Quando um novo produto é criado na função CriarProduto() do controlador o método Preparar() é chamado garantindo que apenas os produtos válidos sejam inseridos no banco de dados