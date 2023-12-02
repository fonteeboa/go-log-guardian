package services

import (
	"errors"

	pkg "github.com/fonteeBoa/go-log-guardian/pkg/domain"

	"github.com/fonteeBoa/go-log-guardian/internal/database/dbhandler"

	"github.com/fonteeBoa/go-log-guardian/internal/database/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// SaveLog saves a log to the database.
//
// It takes a log interface{} as a parameter.
// Returns an error if there was a problem saving the log.
func SaveLog(log interface{}) error {
	gormDB, mongoClient, err := dbhandler.GetConnection()

	if err != nil {
		return err
	}

	if gormDB == nil && mongoClient == nil {
		return errors.New("no valid database connection provided")
	}

	if gormDB != nil {
		errGorm := insertLogGorm(gormDB, log)
		if errGorm != nil {
			return errGorm
		}
	}

	if mongoClient != nil {
		errMongo := insertLogMongo(mongoClient, log)
		if errMongo != nil {
			return errMongo
		}
	}

	return nil

}

// insertLogGorm inserts a log entry into the database using the provided *gorm.DB connection.
//
// The function takes two parameters:
// - db: a *gorm.DB connection to the database.
// - log: an interface{} representing the log entry to be inserted.
//
// The function returns an error type.
func insertLogGorm(db *gorm.DB, log interface{}) error {
	switch log := log.(type) {
	case pkg.LogBase:
		return dbhandler.InsertBaseLog(db, log)
	case pkg.LogFunction:
		return dbhandler.InsertFunctionLog(db, log)
	case pkg.LogDatabase:
		return dbhandler.InsertDatabaseLog(db, log)
	case pkg.LogRequest:
		return dbhandler.InsertRequestLog(db, log)
	default:
		return nil
	}
}

// insertLogMongo inserts a log into a MongoDB database.
//
// db is the MongoDB client.
// log is the log to be inserted. It can be one of the following types:
// - pkg.LogBase
// - pkg.LogFunction
// - pkg.LogDatabase
// - pkg.LogRequest
//
// Returns an error if the insertion fails.
func insertLogMongo(db *mongo.Client, log interface{}) error {
	switch log := log.(type) {
	case pkg.LogBase:
		return mongodb.InsertBaseLog(db, log)
	case pkg.LogFunction:
		return mongodb.InsertFunctionLog(db, log)
	case pkg.LogDatabase:
		return mongodb.InsertDatabaseLog(db, log)
	case pkg.LogRequest:
		return mongodb.InsertRequestLog(db, log)
	default:
		return nil
	}
}
