package elastic

import (
	"encoding/json"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

// getIndexName retorna o nome do índice baseado na variável de ambiente ELASTIC_INDEX.
// Se a variável não estiver definida, retorna o valor padrão "log-guardian-audit".
func getIndexName() string {
	indexName := os.Getenv("ELASTIC_INDEX")
	if indexName == "" {
		return "log-guardian-audit"
	}
	return indexName
}

// InsertBaseLog inserts a base log into the specified index in ElasticSearch.
//
// Parameters:
// - client: A pointer to an elasticsearch.Client instance.
// - log: The base log to be inserted.
//
// Returns:
// - An error if the operation fails.
func InsertBaseLog(client *elasticsearch.Client, log interface{}) error {
	return insertLog(client, getIndexName(), log)
}

// InsertFunctionLog inserts a function log into the specified index in ElasticSearch.
//
// Parameters:
// - client: A pointer to an elasticsearch.Client instance.
// - log: The function log to be inserted.
//
// Returns:
// - An error if the operation fails.
func InsertFunctionLog(client *elasticsearch.Client, log interface{}) error {
	return insertLog(client, getIndexName(), log)
}

// InsertDatabaseLog inserts a database log into the specified index in ElasticSearch.
//
// Parameters:
// - client: A pointer to an elasticsearch.Client instance.
// - log: The database log to be inserted.
//
// Returns:
// - An error if the operation fails.
func InsertDatabaseLog(client *elasticsearch.Client, log interface{}) error {
	return insertLog(client, getIndexName(), log)
}

// InsertRequestLog inserts a request log into the specified index in ElasticSearch.
//
// Parameters:
// - client: A pointer to an elasticsearch.Client instance.
// - log: The request log to be inserted.
//
// Returns:
// - An error if the operation fails.
func InsertRequestLog(client *elasticsearch.Client, log interface{}) error {
	return insertLog(client, getIndexName(), log)
}

// insertLog is a generic function to insert a log into a specified index.
//
// Parameters:
// - client: A pointer to an elasticsearch.Client instance.
// - index: The index name where the log will be inserted.
// - log: The log data to be inserted.
//
// Returns:
// - An error if the operation fails.
func insertLog(client *elasticsearch.Client, index string, log interface{}) error {
	data, err := json.Marshal(log)
	if err != nil {
		return err
	}

	res, err := client.Index(index, esutil.NewJSONReader(data))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
