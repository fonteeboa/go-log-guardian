/**
 * @file connect.go
 * @brief Este arquivo contém funções para conectar ao banco de dados MongoDB e obter coleções.
 */

package mongodb

import (
	"context"
	"os"

	"github.com/fonteeBoa/go-log-guardian/internal/database/migrate"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
 * @brief Estabelece uma conexão com o banco de dados MongoDB.
 *
 * @return Um ponteiro para um mongo.Client e um erro, se houver.
 */
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

/**
 * @brief Retorna uma coleção do MongoDB com base no cliente e no nome da coleção fornecidos.
 *
 * @param client Um ponteiro para uma instância de mongo.Client.
 * @param collectionName O nome da coleção.
 * @return Um ponteiro para uma instância de mongo.Collection.
 */
func getCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	dbName := migrate.GetDbName()
	db := client.Database(dbName)
	return db.Collection(collectionName)
}
