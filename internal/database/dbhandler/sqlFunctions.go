package dbhandler

import (
	pkg "goLogGuardian/pkg/domain"

	"gorm.io/gorm"
)

// InsertBaseLog inserts a BaseLog into the database.
//
// It takes a *gorm.DB and a pkg.BaseLog as parameters.
// It returns an error.
func InsertBaseLog(db *gorm.DB, log pkg.BaseLog) error {
	return db.Create(&log).Error
}

// InsertFunctionLog inserts a function log into the database.
//
// Parameters:
// - db: The gorm.DB object representing the database connection.
// - log: The pkg.FunctionLog object representing the function log to be inserted.
//
// Returns:
// - error: An error, if any occurred during the database insert operation.
func InsertFunctionLog(db *gorm.DB, log pkg.FunctionLog) error {
	return db.Create(&log).Error
}

// InsertDatabaseLog inserts a database log into the given *gorm.DB instance.
//
// It takes the following parameter(s):
// - db: a pointer to a gorm.DB instance representing the database connection.
// - log: a pkg.DatabaseLog struct representing the log to be inserted.
//
// It returns an error if there was an issue inserting the log.
func InsertDatabaseLog(db *gorm.DB, log pkg.DatabaseLog) error {
	return db.Create(&log).Error
}

// InsertRequestLog inserts a request log into the database.
//
// It takes a *gorm.DB object as the first parameter, which represents the database connection,
// and a pkg.RequestLog object as the second parameter, which contains the request log data.
//
// It returns an error if there was an issue inserting the request log into the database.
func InsertRequestLog(db *gorm.DB, log pkg.RequestLog) error {
	return db.Create(&log).Error
}
