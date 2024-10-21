CREATE DATABASE IF NOT EXISTS easyproducts;
USE easyproducts;

DROP TABLE IF EXISTS produtos;

CREATE TABLE produtos(
    id int auto_increment primary key,
    nome varchar(50) not null,
    custo float(10,2) -- Campo float com até 10 dígitos no total, sendo 2 decimais
    quantidade int
) ENGINE=INNODB;