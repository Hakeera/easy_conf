<?php
include("config.php");

class Database
{
    private $pdo;

    public function __construct()
    {
        $connectionUrl = "mysql:host=" . DB_HOST . ";dbname=" . DB_NAME;
        $this->pdo = new PDO($connectionUrl, DB_USER, DB_PASSWORD);
    }

    public function addClient($name, $number, $cellphoneNumber, $documentNumber, $documentType, $email, $situation, $stateInscription, $cest, $address, $city, $state, $cep, $nickname)
    {
        // Verifying required fields
        if (empty($name) || empty($situation) || empty($nickname)) {
            return false;
        }

        $query = $this->pdo->prepare("INSERT INTO clientes VALUES (null, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)");
        $queryExecutionSuccess = $query->execute(array($name, $number, $cellphoneNumber, $documentNumber, $documentType, $email, $situation, $stateInscription, $cest, $address, $city, $state, $cep, $nickname));

        return $queryExecutionSuccess;
    }
}
