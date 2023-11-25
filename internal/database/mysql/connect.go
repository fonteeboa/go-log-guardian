package mysql

import (
	"fmt"
	"os"

	"goLogGuardian/internal/database/migrate"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect establishes a connection to the MySQL database.
//
// It reads the environment variables from the .env file and constructs the connection string.
// The function returns a *gorm.DB instance and an error.
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
