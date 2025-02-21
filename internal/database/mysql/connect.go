/**
 * @file connect.go
 * @brief Este arquivo contém a função para estabelecer uma conexão com o banco de dados MySQL.
 *
 * Ele lê as variáveis de ambiente do arquivo .env e constrói a string de conexão.
 * A função retorna uma instância de *gorm.DB e um erro.
 */

package mysql

import (
	"fmt"
	"os"

	"github.com/fonteeBoa/go-log-guardian/internal/database/migrate"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * @brief Estabelece uma conexão com o banco de dados MySQL.
 *
 * @details Esta função lê as variáveis de ambiente `MYSQL_HOST`, `MYSQL_PORT`, `MYSQL_USER`, `MYSQL_PASSWORD` e `MYSQL_DBNAME`
 * para construir a string de conexão DSN. Em seguida, tenta abrir uma conexão com o banco de dados usando GORM.
 * Se a conexão for bem-sucedida, a função retorna uma instância de *gorm.DB. Caso contrário, retorna um erro.
 * A função também executa migrações no banco de dados.
 *
 * @return *gorm.DB Instância do banco de dados GORM.
 * @return error Erro ocorrido durante a conexão ou migração.
 */
func Connect() (*gorm.DB, error) {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	migrate.MigrateSql(db)

	return db, nil
}
