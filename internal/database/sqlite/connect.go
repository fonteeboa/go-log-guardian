/**
 * @file connect.go
 * @brief This file contains the implementation of the Connect function for establishing a connection to the SQLite database.
 */

package sqlite

import (
	"os"

	"github.com/fonteeBoa/go-log-guardian/internal/database/migrate"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/**
 * @brief Connect establishes a connection to the SQLite database.
 *
 * This function uses the GORM library to open a connection to the SQLite database
 * specified by the environment variable SQLITE_PATH. It also pings the database
 * to ensure the connection is valid and runs any necessary migrations.
 *
 * @return *gorm.DB A pointer to the gorm.DB object representing the database connection.
 * @return error An error object if there was an issue establishing the connection.
 */
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_PATH")), &gorm.Config{})
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
