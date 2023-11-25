package migrate

import (
	"context"
	pkg "goLogGuardian/pkg/domain"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetDbName returns the name of the MongoDB database.
//
// It does not take any parameters.
// It returns a string representing the name of the database.
func GetDbName() string {
	dbName := os.Getenv("MONGODB_DBNAME")

	if dbName == "" {
		dbName = "mydb"
	}
	return dbName
}

// MigrateMongo migrates the MongoDB collections and creates necessary indexes.
//
// It takes a *mongo.Client as a parameter.
// It returns an error if any error occurred during the migration process.
func MigrateMongo(client *mongo.Client) error {
	dbName := GetDbName()
	db := client.Database(dbName)

	collections := []struct {
		Name   string
		Models interface{}
	}{
		{"functionLogs", &pkg.FunctionLog{}},
		{"databaseLogs", &pkg.DatabaseLog{}},
		{"requestLogs", &pkg.RequestLog{}},
		{"baseLogs", &pkg.BaseLog{}},
	}

	for _, col := range collections {
		collection := db.Collection(col.Name)
		if err := collection.Drop(context.Background()); err != nil {
			return err
		}
		if err := db.CreateCollection(context.Background(), col.Name); err != nil {
			return err
		}
		indexes := []mongo.IndexModel{
			// Define os índices desejados para cada coleção, se necessário
		}
		_, err := collection.Indexes().CreateMany(context.Background(), indexes)
		if err != nil {
			return err
		}
	}

	return nil
}

// DoesCollectionExist checks if a collection exists in the MongoDB database.
//
// Parameters:
// - client: The MongoDB client.
// - collectionName: The name of the collection to check.
//
// Returns:
// - bool: true if the collection exists, false otherwise.
func DoesCollectionExist(client *mongo.Client, collectionName string) bool {
	dbName := GetDbName()
	db := client.Database(dbName)

	collections, err := db.ListCollectionNames(context.Background(), bson.M{"name": collectionName})
	if err != nil {
		return false
	}

	for _, col := range collections {
		if col == collectionName {
			return true
		}
	}

	return false
}
