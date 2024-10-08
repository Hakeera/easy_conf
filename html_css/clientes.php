<?php include "../backend/database/Database.php"; ?>

<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cadastro de Clientes</title>
    <link rel="stylesheet" href="style2.css">
</head>

<body>
    <div class="main_grid">
        <div class="logo">
            <h1>EASY</h1>
        </div>

        <div class="ith"><a href="home.html">Home</a></div>
        <div class="page">Controle de Clientes</div>

        <div class="itsel">
            <div class="select">
                <h1>Clientes</h1>
            </div>
            <div class="select">
                <h1>Produtos</h1>
            </div>
            <div class="select">
                <h1>Pedidos</h1>
            </div>
            <div class="select">
                <h1>Contas</h1>
            </div>
            <div class="select">
                <h1>Fornecimento</h1>
            </div>
            <div class="select">
                <h1>Fiscal</h1>
            </div>
        </div>

        <!-- Cadastro de Clientes -->
        <div class="form-container it_vis">
            <h3>Eu quero:</h3>

            <select id="form-container-select">
                <option>Cadastrar um novo cliente</option>
                <option>Atualizar os dados de um cliente já existente</option>
            </select>

            <form method="post" action="client_script.php" id="novo_cliente">
                <div class="button-group">
                    <button id="novo_cliente_btn" type="submit" name="action" value="new_client">Novo Cliente</button>
                </div>

                <label for="nome">Nome:</label>
                <input type="text" id="nome" name="nome" required />

                <label for="telefone">Telefone:</label>
                <input type="text" id="telefone" name="telefone" />

                <label for="celular">Celular</label>
                <input type="text" name="celular" id="celular" />

                <label for="numero_documento">Número do documento</label>
                <input type="text" name="numero_documento" id="numero_documento" />

                <label for="tipo_documento">Tipo do documento:</label>
                <select name="tipo_documento" id="tipo_documento">
                    <option value="CPF">CPF</option>
                    <option value="CNPJ">CNPJ</option>
                </select>
                <br />
                <br />

                <label for="email">E-mail:</label>
                <input type="email" name="email" id="email">

                <label for="inscricao_estadual">Inscrição Estadual:</label>
                <input type="text" name="inscricao_estadual" id="inscricao_estadual">

                <label for="cest">CEST:</label>
                <input type="text" name="cest" id="cest">

                <label for="endereco">Endereço:</label>
                <input type="text" name="endereco" id="endereco">

                <label for="cidade">Cidade:</label>
                <input type="text" id="cidade" name="cidade" />

                <label for="estado">Estado:</label>
                <input type="text" id="estado" name="estado" />

                <label for="cep">CEP:</label>
                <input type="text" name="cep" id=cep">

                <label for="apelido">Apelido:</label>
                <input type="text" name="apelido" id="apelido" required>

                <label for="situacao">Situação:</label>
                <input type="text" id="situacao" name="situacao" required />
            </form>

            <form method="post" action="client_script.php" id="atualizar_cliente">
                <div class="button-group">
                    <button id="novo_cliente_btn" type="submit" name="action" value="new_client">Atualizar cliente</button>
                </div>

                <label for="nome">Nome:</label>
                <input type="text" id="nome-cliente" name="nome" required />

                <label for="telefone">Telefone:</label>
                <input type="text" id="telefone" name="telefone" />

                <label for="celular">Celular</label>
                <input type="text" name="celular" id="celular" />

                <label for="numero_documento">Número do documento</label>
                <input type="text" name="numero_documento" id="numero_documento" />

                <label for="tipo_documento">Tipo do documento:</label>
                <select name="tipo_documento" id="tipo_documento">
                    <option value="CPF">CPF</option>
                    <option value="CNPJ">CNPJ</option>
                </select>
                <br />
                <br />

                <label for="email">E-mail:</label>
                <input type="email" name="email" id="email">

                <label for="inscricao_estadual">Inscrição Estadual:</label>
                <input type="text" name="inscricao_estadual" id="inscricao_estadual">

                <label for="cest">CEST:</label>
                <input type="text" name="cest" id="cest">

                <label for="endereco">Endereço:</label>
                <input type="text" name="endereco" id="endereco">

                <label for="cidade">Cidade:</label>
                <input type="text" id="cidade-cliente" name="cidade" />

                <label for="estado">Estado:</label>
                <input type="text" id="estado-cliente" name="estado" />

                <label for="cep">CEP:</label>
                <input type="text" name="cep" id=cep">

                <label for="apelido">Apelido:</label>
                <input type="text" name="apelido" id="apelido" required>

                <label for="situacao">Situação:</label>
                <input type="text" id="situacao" name="situacao" required />
            </form>
        </div>

        <!-- Lista de Clientes -->
        <div class="client-list">
            <h2>Lista de Clientes</h2>

            <!-- Campo de Busca-->
            <div class="search-container">
                <input type="text" id="search-input" placeholder="Buscar cliente..." onkeyup="searchClient()">
            </div>

            <div class="table-wrapper">
                <table id="clients-table">
                    <thead>
                        <tr>
                            <th>Nome</th>
                            <th>Cidade</th>
                            <th>Estado</th>
                        </tr>
                    </thead>

                    <tbody>
                        <?php
                        $database = new Database();

                        $clients = $database->getClients();

                        foreach ($clients as $client):

                        ?>

                            <tr>
                                <td><?php echo $client["nome"]; ?></td>
                                <td><?php echo $client["cidade"]; ?></td>
                                <td><?php echo $client["estado"]; ?></td>
                            </tr>


                        <?php endforeach; ?>
                    </tbody>
                </table>
            </div>
        </div>

    </div>
</body>
<script src="script.js"></script>

</html>