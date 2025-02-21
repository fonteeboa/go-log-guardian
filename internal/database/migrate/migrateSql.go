package migrate

import (
	pkg "github.com/fonteeBoa/go-log-guardian/pkg/domain"

	"gorm.io/gorm"
)

/**
 * @brief Migra as tabelas SQL.
 *
 * @param db Conexão com o banco de dados gorm.DB.
 * @return Retorna um erro se houve um problema durante a migração.
 */
func MigrateSql(db *gorm.DB) error {
	if err := db.AutoMigrate(&pkg.LogDetails{}, &pkg.LogFunction{}, &pkg.LogDatabase{}, &pkg.LogRequest{}); err != nil {
		return err
	}
	return nil
}

/**
 * @brief Verifica se uma tabela existe no banco de dados.
 *
 * @param db Ponteiro para um objeto gorm.DB representando a conexão com o banco de dados.
 * @param tableName Uma string representando o nome da tabela a ser verificada.
 * @return Retorna um valor booleano indicando se a tabela existe ou não.
 */
func DoesTableExist(db *gorm.DB, tableName string) bool {
	return db.Migrator().HasTable(tableName)
}
