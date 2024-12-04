package dbhandler

import (
	"errors"
	"fmt"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"

	"github.com/fonteeBoa/go-log-guardian/internal/database/elastic"
	"github.com/fonteeBoa/go-log-guardian/internal/database/mongodb"
	"github.com/fonteeBoa/go-log-guardian/internal/database/mysql"
	"github.com/fonteeBoa/go-log-guardian/internal/database/postgres"
	"github.com/fonteeBoa/go-log-guardian/internal/database/sqlite"
)

// GetConnection retrieves the database connection based on the DATABASE_TYPE environment variable.
//
// It returns the *gorm.DB, *mongo.Client, *elasticsearch.Client connections, and an error.
func GetConnection() (*gorm.DB, *mongo.Client, *elasticsearch.Client, error) {
	dbType := os.Getenv("DATABASE_TYPE")
	if dbType == "" {
		fmt.Println("DATABASE_TYPE not set")
		return nil, nil, nil, errors.New("database type not specified")
	}

	var gormDB *gorm.DB
	var mongoClient *mongo.Client
	var elasticClient *elasticsearch.Client
	var err error

	switch dbType {
	case "postgres":
		gormDB, err = postgres.Connect()
		if err != nil {
			return nil, nil, nil, err
		}
	case "mysql":
		gormDB, err = mysql.Connect()
		if err != nil {
			return nil, nil, nil, err
		}
	case "sqlite":
		gormDB, err = sqlite.Connect()
		if err != nil {
			return nil, nil, nil, err
		}
	case "mongodb":
		mongoClient, err = mongodb.Connect()
		if err != nil {
			return nil, nil, nil, err
		}
	case "elastic":
		elasticClient, err = elastic.Connect()
		if err != nil {
			return nil, nil, nil, err
		}
	default:
		return nil, nil, nil, errors.New("unsupported database type")
	}

	return gormDB, mongoClient, elasticClient, nil
}
