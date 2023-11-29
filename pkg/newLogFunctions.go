package pkg

import (
	"time"

	"github.com/FonteeBoa/go-log-guardian/pkg/domain"
)

// newBaseLog creates a new instance of the BaseLog struct.
//
// Parameters:
// - priority: The priority of the log.
// - genericErrMsg: The generic error message.
// - errMsg: The specific error message.
//
// Return type:
// - domain.BaseLog: The newly created BaseLog instance.
func newBaseLog(priority domain.Priority, genericErrMsg string, errMsg string) domain.BaseLog {
	return domain.BaseLog{
		Priority:            priority,
		LogLevel:            domain.PriorityToString[priority],
		Timestamp:           time.Now(),
		GenericErrorMessage: genericErrMsg,
		ErrorMessage:        errMsg,
	}
}

// newFunctionLog creates a new function log with the given priority, function name, generic error message, and error message.
//
// Parameters:
// - priority: the priority of the log
// - functionName: the name of the function
// - genericErrMsg: the generic error message
// - errMsg: the specific error message
//
// Returns:
// - a FunctionLog struct representing the function log
func newFunctionLog(priority domain.Priority, functionName string, genericErrMsg string, errMsg string) domain.FunctionLog {
	return domain.FunctionLog{
		BaseLog: domain.BaseLog{
			Priority:            priority,
			LogLevel:            domain.PriorityToString[priority],
			Timestamp:           time.Now(),
			GenericErrorMessage: genericErrMsg,
			ErrorMessage:        errMsg,
		},
		FunctionName: functionName,
	}
}

// newDatabaseLog creates a new database log entry.
//
// Parameters:
//   - priority: the priority of the log entry (domain.Priority).
//   - tableName: the name of the table (string).
//   - query: the query that was executed (string).
//   - genericErrMsg: the generic error message (string).
//   - errMsg: the specific error message (string).
//
// Returns:
//   - domain.DatabaseLog: the created database log entry.
func newDatabaseLog(priority domain.Priority, tableName string, query string, genericErrMsg string, errMsg string) domain.DatabaseLog {
	return domain.DatabaseLog{
		BaseLog: domain.BaseLog{
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

// newRequestLog creates a new request log with the given parameters.
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
// Return type: domain.RequestLog
func newRequestLog(priority domain.Priority, method string, statusCode int, path string, responseSize int, genericErrMsg string, errMsg string) domain.RequestLog {
	return domain.RequestLog{
		BaseLog: domain.BaseLog{
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
