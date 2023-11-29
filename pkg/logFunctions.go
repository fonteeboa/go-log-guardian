package pkg

import "GoLogGuardian/pkg/domain"

// function is a Go function that takes in a priority, functionName, genericErrMsg, and errMsg as parameters.
//
// It calls the debug function with the provided priority, genericErrMsg, and errMsg arguments.
// Then it returns a new FunctionLog struct by calling the newFunctionLog function with the provided priority, functionName, genericErrMsg, and errMsg arguments.
func function(priority domain.Priority, functionName string, genericErrMsg string, errMsg string) domain.FunctionLog {
	debug(priority, genericErrMsg, errMsg)
	return newFunctionLog(priority, functionName, genericErrMsg, errMsg)
}

// database generates a new database log and returns it.
//
// Parameters:
// - priority: the priority of the log.
// - tableName: the name of the table to log to.
// - query: the query to execute.
// - genericErrMsg: the generic error message.
// - errMsg: the specific error message.
//
// Returns:
// - domain.DatabaseLog: the generated database log.
func database(priority domain.Priority, tableName string, query string, genericErrMsg string, errMsg string) domain.DatabaseLog {
	debug(priority, genericErrMsg, errMsg)
	return newDatabaseLog(priority, tableName, query, genericErrMsg, errMsg)
}

// request generates a new request log with the provided parameters.
//
// Parameters:
//   - priority: The priority of the request.
//   - method: The HTTP method of the request.
//   - statusCode: The HTTP status code of the response.
//   - path: The path of the request.
//   - responseSize: The size of the response in bytes.
//   - genericErrMsg: The generic error message.
//   - errMsg: The specific error message.
//
// Returns:
//   - domain.RequestLog: The generated request log.
func request(priority domain.Priority, method string, statusCode int, path string, responseSize int, genericErrMsg string, errMsg string) domain.RequestLog {
	debug(priority, genericErrMsg, errMsg)
	return newRequestLog(priority, method, statusCode, path, responseSize, genericErrMsg, errMsg)
}

// error logs an error message with the given priority, generic error message, and error message.
//
// priority: The priority of the error message.
// genericErrMsg: The generic error message.
// errMsg: The specific error message.
// Returns a domain.BaseLog.
func error(priority domain.Priority, genericErrMsg string, errMsg string) domain.BaseLog {
	debug(priority, genericErrMsg, errMsg)
	return newBaseLog(priority, genericErrMsg, errMsg)
}
