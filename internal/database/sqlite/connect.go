package sqlite

import (
	"os"

	"github.com/FonteeBoa/go-log-guardian/internal/database/migrate"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Connect establishes a connection to the SQLite database.
//
// It returns a pointer to the gorm.DB object and an error, if any.
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
