package dbhandler

import (
	pkg "github.com/fonteeBoa/go-log-guardian/pkg/domain"

	"gorm.io/gorm"
)

/**
 * @brief Insere um LogDetails no banco de dados.
 *
 * @param db Ponteiro para um objeto gorm.DB representando a conexão com o banco de dados.
 * @param log Objeto pkg.LogDetails representando o log a ser inserido.
 * @return Retorna um erro se houve um problema durante a inserção no banco de dados.
 */
func InsertBaseLog(db *gorm.DB, log pkg.LogDetails) error {
	return db.Create(&log).Error
}

/**
 * @brief Insere um log de função no banco de dados.
 *
 * @param db Objeto gorm.DB representando a conexão com o banco de dados.
 * @param log Objeto pkg.LogFunction representando o log de função a ser inserido.
 * @return Retorna um erro se houve um problema durante a inserção no banco de dados.
 */
func InsertFunctionLog(db *gorm.DB, log pkg.LogFunction) error {
	return db.Create(&log).Error
}

/**
 * @brief Insere um log de banco de dados no banco de dados.
 *
 * @param db Ponteiro para um objeto gorm.DB representando a conexão com o banco de dados.
 * @param log Estrutura pkg.LogDatabase representando o log a ser inserido.
 * @return Retorna um erro se houve um problema durante a inserção no banco de dados.
 */
func InsertDatabaseLog(db *gorm.DB, log pkg.LogDatabase) error {
	return db.Create(&log).Error
}

/**
 * @brief Insere um log de requisição no banco de dados.
 *
 * @param db Objeto gorm.DB representando a conexão com o banco de dados.
 * @param log Objeto pkg.LogRequest contendo os dados do log de requisição.
 * @return Retorna um erro se houve um problema durante a inserção no banco de dados.
 */
func InsertRequestLog(db *gorm.DB, log pkg.LogRequest) error {
	return db.Create(&log).Error
}
