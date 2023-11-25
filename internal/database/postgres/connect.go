package postgres

import (
	"fmt"
	"goLogGuardian/internal/database/migrate"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect establishes a connection to the PostgreSQL database.
//
// It reads the environment variables from the .env file and constructs the connection string.
// The function returns a *gorm.DB instance and an error.
func Connect() (*gorm.DB, error) {
	// Leitura das variáveis de ambiente do arquivo .env
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_EXTERNAL_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	if host == "" {
		host = "127.0.0.1"
	}

	// Construa a string de conexão
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Testa a conexão com o banco de dados
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
