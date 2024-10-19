package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/tabela", func(w http.ResponseWriter, r *http.Request) {
        // Lógica para conectar ao banco de dados ou retornar dados
        fmt.Fprintln(w, "Dados conectados")
    })

    fmt.Println("Conection rodando na porta 8080...")
    http.ListenAndServe(":8080", nil)
}
