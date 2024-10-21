package modelos

//Atributos dos produtos
type Produto struct {
	ID uint64 `json:"id,omitempty"`
	Nome string `json:"nome,omitempty"`
	Custo float64 `json:"custo,omitempty"`
	Quantidade int `json:"quantidade,omitempty"`
}