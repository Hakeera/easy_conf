# easy_conf
Sistema de Gestão - Confecção


# Banco de Dados de Produtos
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


