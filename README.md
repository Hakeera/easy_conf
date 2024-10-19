# easy_conf
Sistema de Gestão - Confecção

# Criando Microserviços e Estabelecendo Conexões

## Estrutura de Pastas

- **Conection/**: Contém o código responsável pela **conexão com o banco de dados** e pela **configuração do servidor**.
- **Table/**: Contém códigos para manipualação e renderização de **tabelas** e **grades**

## Conection/main.go

Este arquivo define um servidor HTTP que escuta na porta 8080 e tem um único manipulador de rotas.

### Código

```go
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

```
### Pacote main: Inicia a aplicação.
#### Função main:
  - Usa http.HandleFunc para definir um manipulador para a rota /tabela, que retorna uma mensagem indicando que os dados estão conectados.
  - Inicia o servidor na porta 8080 e imprime uma mensagem no console quando está em execução.

## Table/rendertable.go
Este arquivo define um servidor HTTP que renderiza uma tabela HTML com dados de exemplo.

### Código

```go
Copiar código
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
```
### Explicação
#### Pacote main: Inicia a aplicação.
- Estrutura Item: Define a estrutura que representa um item da tabela com ID, Nome e Custo.
- Dados de exemplo: Cria uma lista de itens (dummy data) para exibir na tabela.
- Função rendertable:
  - Cria um template HTML para renderizar uma tabela com os itens.
  - Usa html/template para garantir que os dados sejam exibidos corretamente em HTML.
- Função main:
  - Configura o manipulador de rota /tabela para chamar a função renderTable.
  - Inicia o servidor na porta 8080.
### Observações
- Execução: Para rodar o projeto, certifique-se de que o servidor esteja ativo e acesse http://localhost:8080/tabela para visualizar a tabela renderizada.
- Erros: Verifique as mensagens de erro no console para diagnóstico caso algo não funcione conforme esperado.

# Estruturas dos **Bancos de Dados**
## Produtos Geral

| **Formatação**        | **Tipo**                    | **Descrição**                              |
|-----------------------|-----------------------------|--------------------------------------------|
| Produto ID            | Long/String                 | Primary Key                                |
| Modelo                | String                      | Aberto                                     |
| Tecido                | String                      | Aberto                                     |
| Cor                   | String                      | Aberto                                     |
| Cliente               | String                      | Vinculado (Clientes DB)                    |
| Tamanho               | String                      | Semi Aberto (Grades Pré Definidas)          |
| Linha                 | String                      | Aberto                                     |
| Situação              | String                      | Aberto                                     |
| Custo                 | Currency                    | Vinculado (Soma Custos - Custos DB)        |
| Venda                 | Currency                    | Vinculado (Custo * MarkUp - Custos DB)     |
| Descrição             | String                      | Aberto                                     |
| Nome                  | String                      | Vinculado (Modelo + Tecido + Cor + Cliente + Tamanho) |

## Produtos Custos
| **Formatação**        | **Tipo**                    | **Descrição**                              |
|-----------------------|-----------------------------|--------------------------------------------|
| Produto ID            | Long/String                 | Primary Key (Relacional)                   |
| Tecido                | Currency                    | Aberto                                     |
| Aviamento             | Currency                    | Aberto                                     |
| Corte                 | Currency                    | Aberto                                     |
| Costura               | Currency                    | Aberto                                     |
| Transporte            | Currency                    | Aberto                                     |
| Embalagem             | Currency                    | Aberto                                     |
| Estampa/Bordado       | Currency                    | Aberto                                     |
| MarkUp                | Currency                    | Aberto                                     |


# Banco de Dados de Clientes:

| Campo              | Obrigatoriedade |
|--------------------|-----------------|
| Cliente ID         | Obrigatório      |
| Nome               | Obrigatório      |
| Telefone           | -               |
| Celular            | -               |
| Número do Documento| -               |
| CPF/CNPJ           | -               |
| Email              | -               |
| Situação           | Obrigatório      |
| Incricao Estadual  | -               |
| CEST               | -               |
| Endereço           | -               |
| Cidade             | -               |
| Estado             | -               |
| CEP                | -               |
| Apelido            | Obrigatório      |


