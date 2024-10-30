package modelos

import (
	"errors"
	"strings"
)

//Atributos dos produtos
type Produto struct {
	ID uint64 `json:"id,omitempty"`
	Nome string `json:"nome,omitempty"`
	Custo float64 `json:"custo,omitempty"`
	Quantidade int `json:"quantidade,omitempty"`
}

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