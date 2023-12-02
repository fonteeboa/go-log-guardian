package dbhandler

import (
	pkg "github.com/fonteeBoa/go-log-guardian/pkg/domain"

	"gorm.io/gorm"
)

// InsertBaseLog inserts a LogDetails into the database.
//
// It takes a *gorm.DB and a pkg.LogDetails as parameters.
// It returns an error.
func InsertBaseLog(db *gorm.DB, log pkg.LogDetails) error {
	return db.Create(&log).Error
}

// InsertFunctionLog inserts a function log into the database.
//
// Parameters:
// - db: The gorm.DB object representing the database connection.
// - log: The pkg.LogFunction object representing the function log to be inserted.
//
// Returns:
// - error: An error, if any occurred during the database insert operation.
func InsertFunctionLog(db *gorm.DB, log pkg.LogFunction) error {
	return db.Create(&log).Error
}

// InsertDatabaseLog inserts a database log into the given *gorm.DB instance.
//
// It takes the following parameter(s):
// - db: a pointer to a gorm.DB instance representing the database connection.
// - log: a pkg.LogDatabase struct representing the log to be inserted.
//
// It returns an error if there was an issue inserting the log.
func InsertDatabaseLog(db *gorm.DB, log pkg.LogDatabase) error {
	return db.Create(&log).Error
}

// InsertRequestLog inserts a request log into the database.
//
// It takes a *gorm.DB object as the first parameter, which represents the database connection,
// and a pkg.LogRequest object as the second parameter, which contains the request log data.
//
// It returns an error if there was an issue inserting the request log into the database.
func InsertRequestLog(db *gorm.DB, log pkg.LogRequest) error {
	return db.Create(&log).Error
}
