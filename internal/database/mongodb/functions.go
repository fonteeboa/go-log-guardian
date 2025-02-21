/**
 * @file functions.go
 * @brief Funções para inserir logs em coleções MongoDB.
 */

package mongodb

import (
	"context"

	pkg "github.com/fonteeBoa/go-log-guardian/pkg/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

/**
 * @brief Insere um log base na coleção especificada do MongoDB.
 *
 * @param Client Cliente MongoDB usado para conectar ao banco de dados.
 * @param log Log base a ser inserido.
 * @return error Retorna um erro se a inserção falhar.
 */
func InsertBaseLog(Client *mongo.Client, log pkg.LogDetails) error {
	collection := getCollection(Client, "baseLogs")

	_, err := collection.InsertOne(context.Background(), log)
	return err
}

/**
 * @brief Insere um log de função na coleção "functionLogs".
 *
 * @param Client Cliente MongoDB usado para conectar ao banco de dados.
 * @param log Log de função a ser inserido.
 * @return error Retorna um erro se a inserção falhar.
 */
func InsertFunctionLog(Client *mongo.Client, log pkg.LogFunction) error {
	collection := getCollection(Client, "functionLogs")
	_, err := collection.InsertOne(context.Background(), log)
	return err
}

/**
 * @brief Insere um log de banco de dados na coleção especificada.
 *
 * @param Client Cliente MongoDB usado para conectar ao banco de dados.
 * @param log Log de banco de dados a ser inserido.
 * @return error Retorna um erro se a inserção falhar.
 */
func InsertDatabaseLog(Client *mongo.Client, log pkg.LogDatabase) error {
	collection := getCollection(Client, "databaseLogs")
	_, err := collection.InsertOne(context.Background(), log)
	return err
}

/**
 * @brief Insere um log de requisição na coleção "requestLogs" no banco de dados MongoDB.
 *
 * @param Client Ponteiro para um objeto mongo.Client representando o cliente MongoDB.
 * @param log Log de requisição a ser inserido na coleção.
 * @return error Retorna um erro se a inserção falhar.
 */
func InsertRequestLog(Client *mongo.Client, log pkg.LogRequest) error {
	collection := getCollection(Client, "requestLogs")
	_, err := collection.InsertOne(context.Background(), log)
	return err
}
