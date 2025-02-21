/**
 * @file savelogs.go
 * @brief Este arquivo contém funções para salvar logs em diferentes bancos de dados.
 */

package services

import (
	"errors"

	pkg "github.com/fonteeBoa/go-log-guardian/pkg/domain"

	"github.com/fonteeBoa/go-log-guardian/internal/database/dbhandler"

	"github.com/fonteeBoa/go-log-guardian/internal/database/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

/**
 * @brief Salva um log no banco de dados.
 *
 * @param log O log a ser salvo. Pode ser um dos seguintes tipos:
 * - pkg.LogDetails
 * - pkg.LogFunction
 * - pkg.LogDatabase
 * - pkg.LogRequest
 *
 * @return Retorna um erro se houver um problema ao salvar o log.
 */
func SaveLog(log interface{}) error {
	gormDB, mongoClient, err := dbhandler.GetConnection()

	if err != nil {
		return err
	}

	if gormDB == nil && mongoClient == nil {
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

	return nil
}

/**
 * @brief Insere uma entrada de log no banco de dados usando a conexão *gorm.DB fornecida.
 *
 * @param db Conexão *gorm.DB com o banco de dados.
 * @param log O log a ser inserido. Pode ser um dos seguintes tipos:
 * - pkg.LogDetails
 * - pkg.LogFunction
 * - pkg.LogDatabase
 * - pkg.LogRequest
 *
 * @return Retorna um erro se a inserção falhar.
 */
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

/**
 * @brief Insere um log em um banco de dados MongoDB.
 *
 * @param db Cliente MongoDB.
 * @param log O log a ser inserido. Pode ser um dos seguintes tipos:
 * - pkg.LogDetails
 * - pkg.LogFunction
 * - pkg.LogDatabase
 * - pkg.LogRequest
 *
 * @return Retorna um erro se a inserção falhar.
 */
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
