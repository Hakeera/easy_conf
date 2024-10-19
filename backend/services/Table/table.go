package main

import (
    "fmt"

)

// Estrutura para representar uma célula da tabela (celular booleana)
type TableBool struct {
    rows int         // Número de linhas
    cols int         // Número de colunas
    data [][]bool    // Dados armazenados na tabela (valores booleanos)
}

// Função para criar uma nova tabela booleana
func NewTableBool(rows, cols int) *TableBool {
    table := &TableBool{
        rows: rows,
        cols: cols,
        data: make([][]bool, rows),
    }

    for i := range table.data {
        table.data[i] = make([]bool, cols)  // Inicializa cada linha com valores booleanos padrão (false)
    }
    return table
}

// Método para obter os valores da tabela e realizar operações com eles
func (t *TableBool) ObterValores() {
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
func (t *TableBool) SetValor(row, col int, valor bool) {
    if row < t.rows && col < t.cols {
        t.data[row][col] = valor
    }
}

func main() {
    // Cria uma Table de booleanos 3x3
    table := NewTableBool(3, 3)

    // Define alguns valores de exemplo na tabela
    table.SetValor(0, 0, true)
    table.SetValor(1, 1, true)
    table.SetValor(2, 2, true)

    // Percorre a tabela e imprime os valores
    table.ObterValores()
}
