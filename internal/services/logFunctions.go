/**
 * @file logFunctions.go
 * @brief Este arquivo contém funções para gerar logs de diferentes tipos.
 */

package services

import "github.com/fonteeBoa/go-log-guardian/pkg/domain"

/**
 * @brief Gera um log de função.
 *
 * Esta função chama a função Debug com os argumentos fornecidos e retorna uma nova estrutura LogFunction
 * chamando a função NewFunctionLog com os argumentos fornecidos.
 *
 * @param priority A prioridade do log.
 * @param functionName O nome da função.
 * @param genericErrMsg A mensagem de erro genérica.
 * @param errMsg A mensagem de erro específica.
 * @return domain.LogFunction O log de função gerado.
 */
func Function(priority domain.Priority, functionName string, genericErrMsg string, errMsg string) domain.LogFunction {
	Debug(priority, genericErrMsg, errMsg)
	return NewFunctionLog(priority, functionName, genericErrMsg, errMsg)
}

/**
 * @brief Gera um log de banco de dados.
 *
 * Esta função chama a função Debug com os argumentos fornecidos e retorna um novo log de banco de dados
 * chamando a função NewDatabaseLog com os argumentos fornecidos.
 *
 * @param priority A prioridade do log.
 * @param tableName O nome da tabela.
 * @param query A consulta a ser executada.
 * @param genericErrMsg A mensagem de erro genérica.
 * @param errMsg A mensagem de erro específica.
 * @return domain.LogDatabase O log de banco de dados gerado.
 */
func Database(priority domain.Priority, tableName string, query string, genericErrMsg string, errMsg string) domain.LogDatabase {
	Debug(priority, genericErrMsg, errMsg)
	return NewDatabaseLog(priority, tableName, query, genericErrMsg, errMsg)
}

/**
 * @brief Gera um log de requisição.
 *
 * Esta função chama a função Debug com os argumentos fornecidos e retorna um novo log de requisição
 * chamando a função NewRequestLog com os argumentos fornecidos.
 *
 * @param priority A prioridade da requisição.
 * @param method O método HTTP da requisição.
 * @param statusCode O código de status HTTP da resposta.
 * @param path O caminho da requisição.
 * @param responseSize O tamanho da resposta em bytes.
 * @param genericErrMsg A mensagem de erro genérica.
 * @param errMsg A mensagem de erro específica.
 * @return domain.LogRequest O log de requisição gerado.
 */
func Request(priority domain.Priority, method string, statusCode int, path string, responseSize int, genericErrMsg string, errMsg string) domain.LogRequest {
	Debug(priority, genericErrMsg, errMsg)
	return NewRequestLog(priority, method, statusCode, path, responseSize, genericErrMsg, errMsg)
}

/**
 * @brief Gera detalhes do log.
 *
 * Esta função chama a função Debug com os argumentos fornecidos e retorna os detalhes do log
 * gerados pela função NewLogDetails.
 *
 * @param priority O nível de prioridade dos detalhes do log.
 * @param genericErrMsg A mensagem de erro genérica.
 * @param errMsg A mensagem de erro específica.
 * @return domain.LogDetails Os detalhes do log gerados.
 */
func Details(priority domain.Priority, genericErrMsg string, errMsg string) domain.LogDetails {
	Debug(priority, genericErrMsg, errMsg)
	return NewLogDetails(priority, genericErrMsg, errMsg)
}
