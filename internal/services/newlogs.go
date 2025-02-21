/**
 * @file newlogs.go
 * @brief Este arquivo contém funções para criar novas instâncias de diferentes tipos de logs.
 */

package services

import (
	"time"

	"github.com/fonteeBoa/go-log-guardian/pkg/domain"
)

/**
 * @brief Cria uma nova instância da struct LogDetails.
 *
 * @param priority A prioridade do log.
 * @param genericErrMsg A mensagem de erro genérica.
 * @param errMsg A mensagem de erro específica.
 *
 * @return Retorna uma nova instância de domain.LogDetails.
 */
func NewLogDetails(priority domain.Priority, genericErrMsg string, errMsg string) domain.LogDetails {
	return domain.LogDetails{
		Priority:            priority,
		LogLevel:            domain.PriorityToString[priority],
		Timestamp:           time.Now(),
		GenericErrorMessage: genericErrMsg,
		ErrorMessage:        errMsg,
	}
}

/**
 * @brief Cria um novo log de função com a prioridade, nome da função, mensagem de erro genérica e mensagem de erro específica fornecidos.
 *
 * @param priority A prioridade do log.
 * @param functionName O nome da função.
 * @param genericErrMsg A mensagem de erro genérica.
 * @param errMsg A mensagem de erro específica.
 *
 * @return Retorna uma struct domain.LogFunction representando o log da função.
 */
func NewFunctionLog(priority domain.Priority, functionName string, genericErrMsg string, errMsg string) domain.LogFunction {
	return domain.LogFunction{
		LogDetails: domain.LogDetails{
			Priority:            priority,
			LogLevel:            domain.PriorityToString[priority],
			Timestamp:           time.Now(),
			GenericErrorMessage: genericErrMsg,
			ErrorMessage:        errMsg,
		},
		FunctionName: functionName,
	}
}

/**
 * @brief Cria uma nova entrada de log de banco de dados.
 *
 * @param priority A prioridade da entrada de log (domain.Priority).
 * @param tableName O nome da tabela (string).
 * @param query A consulta que foi executada (string).
 * @param genericErrMsg A mensagem de erro genérica (string).
 * @param errMsg A mensagem de erro específica (string).
 *
 * @return Retorna uma nova instância de domain.LogDatabase.
 */
func NewDatabaseLog(priority domain.Priority, tableName string, query string, genericErrMsg string, errMsg string) domain.LogDatabase {
	return domain.LogDatabase{
		LogDetails: domain.LogDetails{
			Priority:            priority,
			LogLevel:            domain.PriorityToString[priority],
			Timestamp:           time.Now(),
			GenericErrorMessage: genericErrMsg,
			ErrorMessage:        errMsg,
		},
		TableName: tableName,
		Query:     query,
	}
}

/**
 * @brief Cria um novo log de requisição com os parâmetros fornecidos.
 *
 * @param priority A prioridade do log.
 * @param method O método HTTP da requisição.
 * @param statusCode O código de status da resposta.
 * @param path O caminho da requisição.
 * @param responseSize O tamanho da resposta em bytes.
 * @param genericErrMsg A mensagem de erro genérica.
 * @param errMsg A mensagem de erro específica.
 *
 * @return Retorna uma nova instância de domain.LogRequest.
 */
func NewRequestLog(priority domain.Priority, method string, statusCode int, path string, responseSize int, genericErrMsg string, errMsg string) domain.LogRequest {
	return domain.LogRequest{
		LogDetails: domain.LogDetails{
			Priority:            priority,
			LogLevel:            domain.PriorityToString[priority],
			Timestamp:           time.Now(),
			GenericErrorMessage: genericErrMsg,
			ErrorMessage:        errMsg,
		},
		Method:       method,
		StatusCode:   statusCode,
		Path:         path,
		ResponseSize: responseSize,
	}
}
