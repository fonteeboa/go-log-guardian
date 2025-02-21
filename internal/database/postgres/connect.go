/**
 * @file connect.go
 * @brief Este arquivo contém a função para estabelecer uma conexão com o banco de dados PostgreSQL.
 */

package postgres

import (
	"fmt"
	"os"

	"github.com/fonteeBoa/go-log-guardian/internal/database/migrate"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/**
 * @brief Estabelece uma conexão com o banco de dados PostgreSQL.
 *
 * Esta função lê as variáveis de ambiente do arquivo .env e constrói a string de conexão.
 * A função retorna uma instância de *gorm.DB e um erro, se houver.
 *
 * @return *gorm.DB Um ponteiro para a instância gorm.DB.
 * @return error Um erro se a conexão não puder ser estabelecida.
 */
func Connect() (*gorm.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_EXTERNAL_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	if host == "" {
		host = "127.0.0.1"
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
