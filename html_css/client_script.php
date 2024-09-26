<?php

include "../backend/database/Database.php";

if (!($_SERVER["REQUEST_METHOD"] === "POST")) die("Invalid request");

if (isset($_POST["action"]) && $_POST["action"] == "new_client") {
    // Preventing XSS
    $name = htmlentities($_POST["nome"]);
    $city = htmlentities($_POST["cidade"]);
    $state = htmlentities($_POST["estado"]);
    $situation = htmlentities($_POST["situacao"]);
    $nicknamame = $name;

    $database = new Database();

    // TODO: Atualizar os 0's quando o form for atualizado.

    if ($database->addClient($name, 0, 0, 0, 0, 0, $situation, 0, 0, 0, $city, $state, 0, $nicknamame)) {
        print "Cliente adicionado com sucesso!"; // TODO: Decidir o que fazer aqui
    }
}
