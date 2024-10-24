# ezConf
ezConf é um backend para gerenciamento de uma aplicação de controle de produção e vendas (ERP). Ele é construído em Go e utiliza uma arquitetura **monolítica** com suporte para _shutdown gracioso_ e outras funcionalidades escaláveis.

## Estrutura do Projeto
Aqui está a estrutura inicial do projeto:

```bash
ezConf/
├── cmd/
│   └── api/
│       └── main.go         # Arquivo principal do servidor HTTP.
├── internal/
│   └── database/           # Módulo para configuração e integração com o banco de dados.
│   └── server/             # Módulo com as rotas e a lógica do servidor.
│       ├── routes.go       # Define as rotas da aplicação.
│       ├── routes_test.go  # Testes unitários das rotas.
│       └── server.go       # Configuração do servidor e middlewares.
├── .env                    # Variáveis de ambiente para configuração do servidor.
├── .gitignore               # Ignora arquivos e diretórios do versionamento Git.
├── docker-compose.yml       # Configuração do Docker Compose para rodar o projeto em containers.
├── go.mod                   # Arquivo de dependências do projeto Go.
├── go.sum                   # Checksum das dependências do projeto Go.
├── Makefile                 # Arquivo para automação de tarefas.
├── README.md                # Documentação do projeto.
└── .air.toml                # Configuração do Air (hot reloading durante o desenvolvimento).
```
## Funcionalidades
- Servidor HTTP em Go: Usando o pacote net/http para gerenciar requisições.
- Shutdown gracioso: O servidor permite concluir requisições pendentes durante o processo de shutdown, garantindo que o encerramento seja suave.
- Rotas e middlewares: O módulo server encapsula as rotas e middlewares da aplicação.
- Configuração por .env: Parâmetros de configuração sensíveis e variáveis de ambiente são carregados a partir do arquivo .env.

### Arquivo main.go
O arquivo main.go contém o ponto de entrada para o servidor. Ele:

- Inicializa o servidor HTTP.
- Configura um processo de shutdown gracioso.
- Inicia a escuta de requisições na porta especificada.

#### Lógica do ```gracefulShutdown```
  
A função gracefulShutdown permite que o servidor finalize as requisições pendentes antes de desligar, utilizando sinais do sistema operacional (SIGINT e SIGTERM). Após o sinal ser capturado, o servidor recebe 5 segundos para concluir as requisições antes de forçar o desligamento.
