package main

import (
    "html/template"
    "net/http"
)

// Estrutura para representar um item da tabela
type Item struct {
    ID   int
    Name string
    Cost float64
}

// Dados de exemplo (dummy data)
var items = []Item{
    {ID: 1, Name: "Foobar", Cost: 1.23},
    {ID: 2, Name: "Fizzbuzz", Cost: 4.56},
    {ID: 3, Name: "Gizmo", Cost: 78.90},
}

// Função para renderizar a tabela HTML
func renderTable(w http.ResponseWriter, r *http.Request) {
    tmpl := `
    <!DOCTYPE html>
    <html lang="pt-BR">
    <head>
        <meta charset="UTF-8">
        <title>Tabela de Itens</title>
        <style>
            table {
                width: 50%;
                border-collapse: collapse;
                margin: 20px auto;
            }
            th, td {
                border: 1px solid #ddd;
                padding: 8px;
                text-align: center;
            }
            th {
                background-color: #f2f2f2;
            }
        </style>
    </head>
    <body>
        <h2>Tabela de Itens</h2>
        <table>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Cost ($)</th>
            </tr>
            {{range .}}
            <tr>
                <td>{{.ID}}</td>
                <td>{{.Name}}</td>
                <td>{{.Cost}}</td>
            </tr>
            {{end}}
        </table>
    </body>
    </html>
    `

    t, err := template.New("table").Parse(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := t.Execute(w, items); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    http.HandleFunc("/tabela", renderTable)

    // Servidor escutando na porta 8080
    http.ListenAndServe(":8080", nil)
}
