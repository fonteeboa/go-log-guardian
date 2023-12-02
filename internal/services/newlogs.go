package services

import (
	"time"

	"github.com/fonteeBoa/go-log-guardian/pkg/domain"
)

// NewLogDetails creates a new instance of the LogDetails struct.
//
// Parameters:
// - priority: The priority of the log.
// - genericErrMsg: The generic error message.
// - errMsg: The specific error message.
//
// Return type:
// - domain.LogDetails: The newly created LogDetails instance.
func NewLogDetails(priority domain.Priority, genericErrMsg string, errMsg string) domain.LogDetails {
	return domain.LogDetails{
		Priority:            priority,
		LogLevel:            domain.PriorityToString[priority],
		Timestamp:           time.Now(),
		GenericErrorMessage: genericErrMsg,
		ErrorMessage:        errMsg,
	}
}

// NewFunctionLog creates a new function log with the given priority, function name, generic error message, and error message.
//
// Parameters:
// - priority: the priority of the log
// - functionName: the name of the function
// - genericErrMsg: the generic error message
// - errMsg: the specific error message
//
// Returns:
// - a LogFunction struct representing the function log
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

// NewDatabaseLog creates a new database log entry.
//
// Parameters:
//   - priority: the priority of the log entry (domain.Priority).
//   - tableName: the name of the table (string).
//   - query: the query that was executed (string).
//   - genericErrMsg: the generic error message (string).
//   - errMsg: the specific error message (string).
//
// Returns:
//   - domain.LogDatabase: the created database log entry.
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

// NewRequestLog creates a new request log with the given parameters.
//
// Parameters:
// - priority: the priority of the log.
// - method: the HTTP method of the request.
// - statusCode: the status code of the response.
// - path: the path of the request.
// - responseSize: the size of the response in bytes.
// - genericErrMsg: a generic error message.
// - errMsg: a specific error message.
//
// Return type: domain.LogRequest
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
