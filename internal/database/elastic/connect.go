package elastic

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

// Connect establishes a connection to the ElasticSearch cluster.
//
// It reads the ElasticSearch URI from the environment variable `ELASTIC_URI`.
//
// Returns:
// - A pointer to an elasticsearch.Client instance.
// - An error, if the connection fails.
func Connect() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTIC_URI"),
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// Test connection
	res, err := client.Info()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	log.Println("ElasticSearch connection established")
	return client, nil
}
