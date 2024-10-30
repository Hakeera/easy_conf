package main

import "fmt"

// Estrutura para representar uma célula da tabela (celular booleana)
type TabelaBool struct {
    rows int         // Número de linhas
    cols int         // Número de colunas
    data [][]bool    // Dados armazenados na tabela (valores booleanos)
}

// Função para criar uma nova tabela booleana
func NovaTabelaBool(rows, cols int) *TabelaBool {
    tabela := &TabelaBool{
        rows: rows,
        cols: cols,
        data: make([][]bool, rows),
    }

    for i := range tabela.data {
        tabela.data[i] = make([]bool, cols)  // Inicializa cada linha com valores booleanos padrão (false)
    }
    return tabela
}

// Método para obter os valores da tabela e realizar operações com eles
func (t *TabelaBool) ObterValores() {
    for i := 0; i < t.rows; i++ {
        for j := 0; j < t.cols; j++ {
            if t.data[i][j] {
                fmt.Printf("Célula [%d,%d] está marcada como True\n", i, j)
            } else {
                fmt.Printf("Célula [%d,%d] está marcada como False\n", i, j)
            }
        }
    }
}

// Método para definir o valor de uma célula específica
func (t *TabelaBool) SetValor(row, col int, valor bool) {
    if row < t.rows && col < t.cols {
        t.data[row][col] = valor
    }
}

func main() {
    // Cria uma tabela de booleanos 3x3
    tabela := NovaTabelaBool(3, 3)

    // Define alguns valores de exemplo na tabela
    tabela.SetValor(0, 0, true)
    tabela.SetValor(1, 1, true)
    tabela.SetValor(2, 2, true)

    // Percorre a tabela e imprime os valores
    tabela.ObterValores()
}
