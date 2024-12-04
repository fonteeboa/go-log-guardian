package services

import (
	"errors"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/fonteeBoa/go-log-guardian/internal/database/dbhandler"
	"github.com/fonteeBoa/go-log-guardian/internal/database/elastic"
	"github.com/fonteeBoa/go-log-guardian/internal/database/mongodb"
	pkg "github.com/fonteeBoa/go-log-guardian/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// SaveLog saves a log to the database.
//
// It takes a log interface{} as a parameter.
// Returns an error if there was a problem saving the log.
func SaveLog(log interface{}) error {
	gormDB, mongoClient, elasticClient, err := dbhandler.GetConnection()

	if err != nil {
		return err
	}

	if gormDB == nil && mongoClient == nil && elasticClient == nil {
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

	if elasticClient != nil {
		errElastic := insertLogElastic(elasticClient, log)
		if errElastic != nil {
			return errElastic
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
	case pkg.LogDetails:
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
// - pkg.LogDetails
// - pkg.LogFunction
// - pkg.LogDatabase
// - pkg.LogRequest
// Returns:
// Returns an error if the insertion fails.
func insertLogMongo(db *mongo.Client, log interface{}) error {
	switch log := log.(type) {
	case pkg.LogDetails:
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

// insertLogElastic inserts a log into an ElasticSearch index.
//
// Parameters:
// - client: ElasticSearch client instance.
// log is the log to be inserted. It can be one of the following types:
// - pkg.LogDetails
// - pkg.LogFunction
// - pkg.LogDatabase
// - pkg.LogRequest
// Returns:
// - An error if the operation fails.
func insertLogElastic(client *elasticsearch.Client, log interface{}) error {
	switch log := log.(type) {
	case pkg.LogDetails:
		return elastic.InsertBaseLog(client, log)
	case pkg.LogFunction:
		return elastic.InsertFunctionLog(client, log)
	case pkg.LogDatabase:
		return elastic.InsertDatabaseLog(client, log)
	case pkg.LogRequest:
		return elastic.InsertRequestLog(client, log)
	default:
		return nil
	}
}
