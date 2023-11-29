package mongodb

import (
	"GoLogGuardian/internal/database/migrate"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect establishes a connection to the MongoDB database.
//
// It takes no parameters.
// It returns a pointer to a mongo.Client and an error.
func Connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	migrate.MigrateMongo(client)

	return client, nil
}

// getCollection returns a MongoDB collection based on the provided client and collection name.
//
// Parameters:
// - client: A pointer to a mongo.Client instance.
// - collectionName: The name of the collection.
//
// Returns:
// - A pointer to a mongo.Collection instance.
func getCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	dbName := migrate.GetDbName()
	db := client.Database(dbName)
	return db.Collection(collectionName)
}
