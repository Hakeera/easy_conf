package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Abrir conexão com banco de dados e retorna-la
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexãoBanco)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}