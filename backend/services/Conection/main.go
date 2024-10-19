package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/tabela", func(w http.ResponseWriter, r *http.Request) {
        // Lógica para retornar os dados da tabela
        fmt.Fprintln(w, "Dados da Tabela")
    })

    fmt.Println("Servidor rodando na porta 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Erro ao iniciar o servidor:", err)
    }
}
