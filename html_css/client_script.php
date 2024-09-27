<?php

include "../backend/database/Database.php";

if (!($_SERVER["REQUEST_METHOD"] === "POST")) die("Requisicao invalida");

if (isset($_POST["action"]) && $_POST["action"] == "new_client") {
    // Preventing XSS
    $nome = htmlentities($_POST["nome"]);
    $telefone = htmlentities($_POST["telefone"]);
    $celular = htmlentities($_POST["celular"]);
    $numeroDocumento = htmlentities($_POST["numero_documento"]);
    $tipoDocumento = htmlentities($_POST["tipo_documento"]);
    $email = htmlentities($_POST["email"]);
    $inscricaoEstadual = htmlentities($_POST["inscricao_estadual"]);
    $cest = htmlentities($_POST["cest"]);
    $endereco = htmlentities($_POST["endereco"]);
    $cidade = htmlentities($_POST["cidade"]);
    $estado = htmlentities($_POST["estado"]);
    $cep = htmlentities($_POST["cep"]);
    $apelido = htmlentities($_POST["apelido"]);
    $situacao = htmlentities($_POST["situacao"]);

    if ($tipoDocumento != "CPF" && $tipoDocumento != "CNPJ") die("Tipo de documento invalido");

    $database = new Database();

    if ($database->addClient($nome, $telefone, $celular, $numeroDocumento, $tipoDocumento, $email, $situacao, $inscricaoEstadual, $cest, $endereco, $cidade, $estado, $cep, $apelido)) {
        print "Cliente adicionado com sucesso!"; // TODO: Decidir o que fazer aqui
    }
}
