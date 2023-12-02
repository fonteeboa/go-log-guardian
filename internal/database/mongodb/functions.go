package mongodb

import (
	"context"

	pkg "github.com/fonteeBoa/go-log-guardian/pkg/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

// InsertBaseLog inserts a base log into the specified MongoDB collection.
//
// Client is the MongoDB client used to connect to the database.
// log is the base log to be inserted.
// Returns an error if the insertion fails.
func InsertBaseLog(Client *mongo.Client, log pkg.LogDetails) error {
	collection := getCollection(Client, "baseLogs")

	_, err := collection.InsertOne(context.Background(), log)
	return err
}

// InsertFunctionLog inserts a function log into the "functionLogs" collection.
//
// Client is a MongoDB client.
// log is the function log to be inserted.
// Returns an error if the insertion fails.
func InsertFunctionLog(Client *mongo.Client, log pkg.LogFunction) error {
	collection := getCollection(Client, "functionLogs")
	_, err := collection.InsertOne(context.Background(), log)
	return err
}

// InsertDatabaseLog inserts a database log into the specified collection.
//
// It takes a mongo.Client object and a pkg.LogDatabase object as parameters.
// It returns an error if the insertion fails.
func InsertDatabaseLog(Client *mongo.Client, log pkg.LogDatabase) error {
	collection := getCollection(Client, "databaseLogs")
	_, err := collection.InsertOne(context.Background(), log)
	return err
}

// InsertRequestLog inserts a request log into the "requestLogs" collection in the MongoDB database.
//
// Parameters:
// - Client: A pointer to a mongo.Client object representing the MongoDB client.
// - log: The request log to be inserted in the collection.
//
// Returns:
// - error: An error, if any, encountered during the insertion process.
func InsertRequestLog(Client *mongo.Client, log pkg.LogRequest) error {
	collection := getCollection(Client, "requestLogs")
	_, err := collection.InsertOne(context.Background(), log)
	return err
}
