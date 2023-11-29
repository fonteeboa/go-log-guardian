package pkg

import (
	"github.com/FonteeBoa/go-log-guardian/internal/services"
	"github.com/FonteeBoa/go-log-guardian/pkg/domain"
)

// LogFunction logs the function execution and saves the log data if the environment allows it.
//
// Parameters:
// - priority: the priority of the log.
// - functionName: the name of the function being logged.
// - genericErrMsg: the generic error message.
// - errMsg: the specific error message.
//
// Returns:
// - bool: true if the log data was successfully saved, false otherwise.
// - domain.FunctionLog: the logged data.
func LogFunction(priority domain.Priority, functionName string, genericErrMsg string, errMsg string) (bool, domain.FunctionLog) {
	debug(priority, genericErrMsg, errMsg)
	logData := function(priority, functionName, genericErrMsg, errMsg)

	logDb := checkEnvironment()
	if logDb {
		services.SaveLog(logData)
	}

	return true, logData
}

// LogDataBase is a function that logs a database operation and saves the log if the environment allows it.
//
// It takes the following parameters:
// - priority: the priority of the log entry (domain.Priority type).
// - tableName: the name of the table being queried (string type).
// - query: the database query string (string type).
// - genericErrMsg: a generic error message (string type).
// - errMsg: a specific error message (string type).
//
// It returns a boolean value indicating if the log was saved successfully and the logged data (domain.DatabaseLog type).
func LogDataBase(priority domain.Priority, tableName string, query string, genericErrMsg string, errMsg string) (bool, domain.DatabaseLog) {
	debug(priority, genericErrMsg, errMsg)
	logData := database(priority, tableName, query, genericErrMsg, errMsg)

	logDb := checkEnvironment()
	if logDb {
		services.SaveLog(logData)
	}

	return true, logData
}

// LogRequests logs the requests and saves them to the database.
//
// It takes in the following parameters:
//   - priority: the priority of the request
//   - method: the request method
//   - statusCode: the status code of the response
//   - path: the path of the request
//   - responseSize: the size of the response
//   - genericErrMsg: the generic error message
//   - errMsg: the error message
//
// It returns a boolean indicating whether the log was saved successfully and
// a RequestLog struct containing the log data.
func LogRequests(priority domain.Priority, method string, statusCode int, path string, responseSize int, genericErrMsg string, errMsg string) (bool, domain.RequestLog) {
	debug(priority, genericErrMsg, errMsg)
	logData := request(priority, method, statusCode, path, responseSize, genericErrMsg, errMsg)

	logDb := checkEnvironment()
	if logDb {
		services.SaveLog(logData)
	}

	return true, logData
}

// LogError logs an error message with the given priority, generic error message, and specific error message.
//
// It returns a boolean value indicating if the error was logged successfully and the log data.
func LogError(priority domain.Priority, genericErrMsg string, errMsg string) (bool, domain.BaseLog) {
	debug(priority, genericErrMsg, errMsg)
	logData := error(priority, genericErrMsg, errMsg)

	logDb := checkEnvironment()
	if logDb {
		services.SaveLog(logData)
	}

	return true, logData
}
