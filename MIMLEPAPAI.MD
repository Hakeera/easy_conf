main.go:
    Inicia a aplicação, carrega as configurações e inicia o servidor HTTP, escutando na porta definida
        .Carrega as configurações com config.Carregar()
        .Gera o router (definido no pacote router)
        .Inicia o servidor HTTP na porta configurada, utilizando Gorilla Mux para gerenciar as rotas

config/config.go:
    Carrega configurações e define a string de conexão com o banco de dados, bem como a porta da API
        .Lê e carrega variáveis de ambiente com godotenv
        .Define a porta da API e monta a string de conexão com o banco de dados, usada por outras partes do projeto

banco/banco.go
    Gerencia a conexão com o banco de dados, fornecendo a função para abrir uma conexão e verificar sua validade
        .Oferece a função Conectar() para abrir e validar a conexão com o banco de dados MySQL
        .Essa conexão é utilizada pelos controladores e repositórios para executar operações SQL

router/router.go e router/rotas:
    Define as rotas da API e associa as funções especificadas nos controladores. O arquivo rotas.go configura as rotas usando o router do Gorilla Mux
        .Utiliza Gorilla Mux para definir as rotas da API
        .A função Configurar em rotas.go adiciona todas as rotas de produtos e as associa às funções nos controladores

controllers/produtos.go
    Implementa a lógica de negócios para manipulação de produtos(CRUD). Utiliza as funções de respostas do resposta.go e a conexão do banco de dados do arquivo banco.go
        .Implementa funções como CriarProduto, BuscarProdutos, AtualizarProduto e DeletarProduto
        .Utiliza o pacote banco para abrir conexões com o banco de dados e chama os repositórios para manipular dados
        .Envia respostas formatadas em JSON usando as funções do pacote respostas

repositorios/produtos.go
    Organiza a lógia de banco de dados
    Permite que as funções nos controladores sejam mais limpas e focadas na lógica de negócios, as operações de banco de dados é do repositório
        .Encapsula a lógica de interação com a tabela produtos no banco de dados
        .Usado pelos controladores para garantir que toda interação com o banco seja centralizada e organizada

modelos/Produto.go
    Define a estrutura do modelo de dados Produto e implementa métodos de validação e formatação
        .Define a estrutura da entidade Produto com atributos como ID, Nome, Custo e Quantidade
        .Implementa métodos de:
            Validação: Verifica se os atributos do produto são válidos
            Formatação: Remove espaços em branco do nome
            .Usado pelos controladores para manipular dados recebidos ou enviados

respostas/respostas.go
    Formata e envia respostas em JSON, tanto para sucesso quanto para erro
        .Define funções para enviar respostas JSON para o cliente:
            JSON: Envia respostas de sucesso com qualquer dado
            Erro: Envia mensagens de erro em formato JSON
        .Utilizado pelos controladores para padronizar a saída da API

FLUXO DA REQUISIÇÃO:
    1. O cliente (Postman, navegador ou aplicação frontend) faz uma requisição HTTP para o servidor
        Exemplo:
            Metodo: POST
            URL: http://localhost:9000/produtos
            Corpo (JSON):
                {
                    "nome": "Produto",
                    "custo": 20.99,
                    "quantidade": 10
                }
    2. O router recebe a requisição
        No main.go, o servidor é iniciado com a função:
            log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
            .O router foi gerado com router.Gerar()
            .Gerar() é definida no arquivo router/router.go, que chama a função Configurar do pacote rotas:
                func Gerar() *mux.Router {
                    r := mux.NewRouter()
                    return rotas.Configurar(r)
                }
    3. Configuração das rotas
        No arquivo rotas/produtos.go, a rota /produtos é definida:
            var rotasProdutos = []Rota{
                {
                    URI:    "/produtos",
                    Metodo: http.MethodPost,
                    Funcao: controllers.CriarProduto,
                },
            }
        A função Configurar no rotas/rotas.go adiciona essas rotas ao router:
            func Configurar(r *mux.Router) *mux.Router {
                for _, rota := range rotasProdutos {
                    r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
                }
                return r
            }
    4. O router direciona a requisição para o controlador
        Quando a requisição POST /produtos é recebida, o router chama a função CriarProduto no controlador
    5. Controlador: Função CriarProduto
        No arquivo controllers/produtos.go, a função CriarProduto:
            1. Le o corpo da requisição
                corpoRequest, erro := ioutil.ReadAll(r.Body)
            2. Transforma o JSON recebido em um objeto do tipo Produto
                var produto modelos.Produto
                json.Unmarshal(corpoRequest, &produto)
            3. Valida e formata o produto
                produto.Preparar()
            4. Abre uma conexão com o banco de dados
                db, erro := banco.Conectar()
                defer db.Close()
            5. Chama o repositório para criar o produto
                repositorio := repositorios.NovoRepositorioDeProdutos(db)
                produto.ID, erro = repositorio.Criar(produto)
            6. Envia uma resposta JSON ao cliente
                respostas.JSON(w, http.StatusCreated, produto)
    6. Repositório: Função Criar
        No arquivo repositorios/produtos.go, a função Criar executa a inserção no banco de dados
            1. Prepara o comando SQL
                statement, erro := repositorio.db.Prepare(
                "insert into produtos (nome, custo, quantidade) values(?, ?, ?)",
            )
            2. Executa o comando com os dados do produto
                resultado, erro := statement.Exec(produto.Nome, produto.Custo, produto.Quantidade)
            3. Obtém o ID do produto recém-criado
                ultimoIDInserido, erro := resultado.LastInsertId()
            4. Retorna o ID para o controlador
    7. Resposta ao Cliente
        Depois que o produto é criado com sucesso, o controlador envia a seguinte resposta ao cliente
            Resposta (JSON)
            {
                "id": 1,
                "nome": "Produto",
                "custo": 20.99,
                "quantidade": 10
            }