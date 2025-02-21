package migrate

import (
	"context"
	"os"

	pkg "github.com/fonteeBoa/go-log-guardian/pkg/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/**
 * @brief Retorna o nome do banco de dados MongoDB.
 *
 * @return Uma string representando o nome do banco de dados.
 */
func GetDbName() string {
	dbName := os.Getenv("MONGODB_DBNAME")

	if dbName == "" {
		dbName = "mydb"
	}
	return dbName
}

/**
 * @brief Migra as coleções do MongoDB e cria os índices necessários.
 *
 * @param client Ponteiro para um objeto mongo.Client representando a conexão com o MongoDB.
 * @return Retorna um erro se houve um problema durante a migração.
 */
func MigrateMongo(client *mongo.Client) error {
	dbName := GetDbName()
	db := client.Database(dbName)

	collections := []struct {
		Name   string
		Models interface{}
	}{
		{"functionLogs", &pkg.LogFunction{}},
		{"databaseLogs", &pkg.LogDatabase{}},
		{"requestLogs", &pkg.LogRequest{}},
		{"baseLogs", &pkg.LogDetails{}},
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

/**
 * @brief Verifica se uma coleção existe no banco de dados MongoDB.
 *
 * @param client Ponteiro para um objeto mongo.Client representando a conexão com o MongoDB.
 * @param collectionName Uma string representando o nome da coleção a ser verificada.
 * @return Retorna um valor booleano indicando se a coleção existe ou não.
 */
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
