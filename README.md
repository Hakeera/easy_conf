# Projeto API de Produtos

Este projeto é uma API desenvolvida em Go para gerenciamento de produtos, incluindo operações CRUD (Criar, Ler, Atualizar e Deletar). A aplicação utiliza Gorilla Mux para roteamento e MySQL como banco de dados. Abaixo, detalhamos cada componente da aplicação.

## Estrutura do Projeto

### Arquivos Principais

- **`main.go`**: Inicia a aplicação e o servidor HTTP.
  - Carrega as configurações com `config.Carregar()`.
  - Gera o router (definido no pacote `router`).
  - Inicia o servidor HTTP na porta configurada.

- **`config/config.go`**: Carrega as configurações do ambiente.
  - Lê variáveis de ambiente com `godotenv`.
  - Define a porta da API e configura a string de conexão com o banco de dados.

- **`banco/banco.go`**: Gerencia a conexão com o banco de dados MySQL.
  - Função `Conectar()` para abrir e validar a conexão.
  
- **`router/router.go` e `router/rotas`**: Define as rotas da API.
  - Utiliza Gorilla Mux para configurar as rotas.
  - A função `Configurar` em `rotas.go` associa as rotas aos controladores.

### Componentes

- **`controllers/produtos.go`**: Implementa a lógica de negócios para manipulação de produtos.
  - Funções como `CriarProduto`, `BuscarProdutos`, `AtualizarProduto` e `DeletarProduto`.
  - Utiliza o pacote `banco` para conexão e o pacote `respostas` para formatar a resposta.

- **`repositorios/produtos.go`**: Encapsula a lógica de banco de dados da tabela produtos.
  - Centraliza a interação com a tabela de produtos, mantendo os controladores focados na lógica de negócios.

- **`modelos/Produto.go`**: Define o modelo `Produto`.
  - Estrutura com atributos como `ID`, `Nome`, `Custo`, `Quantidade`.
  - Métodos de validação e formatação.

- **`respostas/respostas.go`**: Formata e envia respostas JSON.
  - Funções `JSON` e `Erro` para envio de respostas padronizadas para o cliente.

## Fluxo de uma Requisição

1. **Requisição do Cliente**: O cliente (Postman, navegador, ou app frontend) envia uma requisição HTTP para a API.

    **Exemplo**:
    - Método: `POST`
    - URL: `http://localhost:9000/produtos`
    - Corpo (JSON):
      ```json
      {
        "nome": "Produto",
        "custo": 20.99,
        "quantidade": 10
      }
      ```

2. **Recebimento pelo Router**:
   - O servidor é iniciado em `main.go` com o comando:
     ```go
     log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
     ```
   - `router.Gerar()` gera o router usando a função `Configurar` do pacote `rotas`.

3. **Configuração das Rotas**:
   - No arquivo `rotas/produtos.go`, a rota `/produtos` é definida:
     ```go
     var rotasProdutos = []Rota{
         {
             URI:    "/produtos",
             Metodo: http.MethodPost,
             Funcao: controllers.CriarProduto,
         },
     }
     ```
   - `Configurar` em `rotas/rotas.go` associa essas rotas ao router.

4. **Encaminhamento ao Controlador**:
   - A requisição é direcionada ao controlador que implementa a lógica de manipulação.

5. **Controlador: Função CriarProduto**:
   - Lê o corpo da requisição e transforma o JSON em um objeto `Produto`.
   - Valida e formata o objeto.
   - Abre uma conexão com o banco de dados.
   - Chama o repositório para inserir o produto.
   - Envia uma resposta JSON para o cliente.

6. **Repositório: Função Criar**:
   - Prepara e executa o comando SQL para inserir o produto.
   - Retorna o ID do produto criado para o controlador.

7. **Resposta ao Cliente**:
   - Após a criação do produto, a resposta JSON é enviada:
     ```json
     {
       "id": 1,
       "nome": "Produto",
       "custo": 20.99,
       "quantidade": 10
     }
     ```

## Dependências

- **Go**: Linguagem de programação utilizada.
- **Gorilla Mux**: Biblioteca de roteamento para Go.
- **godotenv**: Carregamento de variáveis de ambiente.
- **MySQL**: Banco de dados relacional.

## Configuração do Ambiente
