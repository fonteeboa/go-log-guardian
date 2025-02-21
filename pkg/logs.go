/**
 * @file logs.go
 * @brief This file contains functions for logging various types of messages and saving them to a database if the environment allows it.
 */

package pkg

import (
	"github.com/fonteeBoa/go-log-guardian/internal/services"
	"github.com/fonteeBoa/go-log-guardian/pkg/domain"
)

/**
 * @brief Logs the function execution and saves the log data if the environment allows it.
 *
 * @param priority The priority of the log.
 * @param functionName The name of the function being logged.
 * @param genericErrMsg The generic error message.
 * @param errMsg The specific error message.
 * @return bool True if the log data was successfully saved, false otherwise.
 * @return domain.LogFunction The logged data.
 */
func LogFunction(priority domain.Priority, functionName string, genericErrMsg string, errMsg string) (bool, domain.LogFunction) {
	services.Debug(priority, genericErrMsg, errMsg)
	logData := services.Function(priority, functionName, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}

	return true, logData
}

/**
 * @brief Logs a services.Database operation and saves the log if the environment allows it.
 *
 * @param priority The priority of the log entry.
 * @param tableName The name of the table being queried.
 * @param query The services.Database query string.
 * @param genericErrMsg A generic error message.
 * @param errMsg A specific error message.
 * @return bool True if the log was saved successfully, false otherwise.
 * @return domain.LogDatabase The logged data.
 */
func LogDataBase(priority domain.Priority, tableName string, query string, genericErrMsg string, errMsg string) (bool, domain.LogDatabase) {
	services.Debug(priority, genericErrMsg, errMsg)
	logData := services.Database(priority, tableName, query, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}

	return true, logData
}

/**
 * @brief Logs the requests and saves them to the services.Database.
 *
 * @param priority The priority of the request.
 * @param method The request method.
 * @param statusCode The status code of the response.
 * @param path The path of the request.
 * @param responseSize The size of the response.
 * @param genericErrMsg The generic error message.
 * @param errMsg The error message.
 * @return bool True if the log was saved successfully, false otherwise.
 * @return domain.LogRequest The logged data.
 */
func LogRequests(priority domain.Priority, method string, statusCode int, path string, responseSize int, genericErrMsg string, errMsg string) (bool, domain.LogRequest) {
	services.Debug(priority, genericErrMsg, errMsg)
	logData := services.Request(priority, method, statusCode, path, responseSize, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}

	return true, logData
}

/**
 * @brief Logs the given error message with the specified priority.
 *
 * @param priority The priority of the log.
 * @param genericErrMsg The generic error message.
 * @param errMsg The specific error message.
 * @return bool True if the log was successfully saved to the database, false otherwise.
 * @return domain.LogDetails The log details.
 */
func Log(priority domain.Priority, genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	services.Debug(priority, genericErrMsg, errMsg)
	logData := services.Details(priority, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}

	return true, logData
}

/**
 * @brief Logs an emergency message.
 *
 * @param genericErrMsg The generic error message.
 * @param errMsg The specific error message.
 * @return bool True if the log was successfully saved, false otherwise.
 * @return domain.LogDetails The log details.
 */
func LogEmerg(genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	logData := services.Details(domain.LOG_EMERG, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}
	return true, logData
}

/**
 * @brief Logs a critical error message.
 *
 * @param genericErrMsg The generic error message.
 * @param errMsg The specific error message.
 * @return bool True if the log was successfully saved, false otherwise.
 * @return domain.LogDetails The log details.
 */
func LogCritical(genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	logData := services.Details(domain.LOG_CRIT, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}
	return true, logData
}

/**
 * @brief Logs an error message.
 *
 * @param priority The priority of the error.
 * @param genericErrMsg The generic error message.
 * @param errMsg The specific error message.
 * @return bool True if the log was successfully saved, false otherwise.
 * @return domain.LogDetails The log details.
 */
func LogError(priority domain.Priority, genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	logData := services.Details(domain.LOG_ERR, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}
	return true, logData
}

/**
 * @brief Generates a log alert and saves it in the database.
 *
 * @param genericErrMsg The generic error message.
 * @param errMsg The specific error message.
 * @return bool True if the log alert was saved in the database, false otherwise.
 * @return domain.LogDetails The details of the generated log alert.
 */
func LogAlert(genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	logData := services.Details(domain.LOG_ALERT, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}
	return true, logData
}

/**
 * @brief Logs a warning message.
 *
 * @param genericErrMsg The generic error message.
 * @param errMsg The specific error message.
 * @return bool True if the log was saved successfully, false otherwise.
 * @return domain.LogDetails The log details.
 */
func LogWarning(genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	logData := services.Details(domain.LOG_WARNING, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}
	return true, logData
}

/**
 * @brief Logs a notice message.
 *
 * @param genericErrMsg The generic error message.
 * @param errMsg The specific error message.
 * @return bool True if the log was saved successfully, false otherwise.
 * @return domain.LogDetails The log details.
 */
func LogNotice(genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	logData := services.Details(domain.LOG_NOTICE, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}
	return true, logData
}

/**
 * @brief Logs information.
 *
 * @param genericErrMsg The generic error message.
 * @param errMsg The specific error message.
 * @return bool True if the log was saved successfully, false otherwise.
 * @return domain.LogDetails The log details.
 */
func LogInfo(genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	logData := services.Details(domain.LOG_INFO, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}
	return true, logData
}

/**
 * @brief Logs a debug message.
 *
 * @param genericErrMsg The generic error message.
 * @param errMsg The specific error message.
 * @return bool True if the log was saved successfully, false otherwise.
 */
func LogDebug(genericErrMsg string, errMsg string) bool {
	services.Debug(domain.LOG_DEBUG, genericErrMsg, errMsg)
	return true
}
