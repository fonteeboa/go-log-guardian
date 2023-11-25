package migrate

import (
	pkg "goLogGuardian/pkg/domain"

	"gorm.io/gorm"
)

// MigrateSql migrates the SQL tables.
//
// db: the gorm.DB connection.
// Returns: an error if there was a problem during migration.
func MigrateSql(db *gorm.DB) error {
	if err := db.AutoMigrate(&pkg.BaseLog{}, &pkg.FunctionLog{}, &pkg.DatabaseLog{}, &pkg.RequestLog{}); err != nil {
		return err
	}
	return nil
}

// DoesTableExist checks if a table exists in the database.
//
// Parameters:
// - db: a pointer to a gorm.DB object representing the database connection.
// - tableName: a string representing the name of the table to check.
//
// Returns:
// - a boolean value indicating whether the table exists or not.
func DoesTableExist(db *gorm.DB, tableName string) bool {
	return db.Migrator().HasTable(tableName)
}
