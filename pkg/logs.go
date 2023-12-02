package pkg

import (
	"github.com/fonteeBoa/go-log-guardian/internal/services"
	"github.com/fonteeBoa/go-log-guardian/pkg/domain"
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
// - domain.LogFunction: the logged data.
func LogFunction(priority domain.Priority, functionName string, genericErrMsg string, errMsg string) (bool, domain.LogFunction) {
	services.Debug(priority, genericErrMsg, errMsg)
	logData := services.Function(priority, functionName, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}

	return true, logData
}

// LogDataBase is a function that logs a services.Database operation and saves the log if the environment allows it.
//
// It takes the following parameters:
// - priority: the priority of the log entry (domain.Priority type).
// - tableName: the name of the table being queried (string type).
// - query: the services.Database query string (string type).
// - genericErrMsg: a generic error message (string type).
// - errMsg: a specific error message (string type).
//
// It returns a boolean value indicating if the log was saved successfully and the logged data (domain.LogDatabase type).
func LogDataBase(priority domain.Priority, tableName string, query string, genericErrMsg string, errMsg string) (bool, domain.LogDatabase) {
	services.Debug(priority, genericErrMsg, errMsg)
	logData := services.Database(priority, tableName, query, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}

	return true, logData
}

// LogRequests logs the requests and saves them to the services.Database.
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
// a LogRequest struct containing the log data.
func LogRequests(priority domain.Priority, method string, statusCode int, path string, responseSize int, genericErrMsg string, errMsg string) (bool, domain.LogRequest) {
	services.Debug(priority, genericErrMsg, errMsg)
	logData := services.Request(priority, method, statusCode, path, responseSize, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}

	return true, logData
}

// Log logs the given error message with the specified priority and returns a boolean value indicating if the log was successfully saved to the database and the log details.
//
// - priority: The priority of the log.
// - genericErrMsg: The generic error message.
// - errMsg: The specific error message.
// Return type(s): bool, domain.LogDetails
func Log(priority domain.Priority, genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	services.Debug(priority, genericErrMsg, errMsg)
	logData := services.Details(priority, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}

	return true, logData
}

// LogEmerg logs an emergency message and returns true along with the log details.
//
// It takes two parameters:
// - genericErrMsg (string): The generic error message.
// - errMsg (string): The specific error message.
//
// It returns a boolean value indicating the success of the log operation and a domain.LogDetails struct containing the log details.
func LogEmerg(genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	logData := services.Details(domain.LOG_EMERG, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}
	return true, logData
}

// LogCritical logs a critical error message along with its details.
//
// It takes two parameters:
// - genericErrMsg: a string representing a generic error message.
// - errMsg: a string representing the specific error message.
//
// It returns a boolean value indicating the success of the logging operation and a LogDetails struct containing the logged details.
func LogCritical(genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	logData := services.Details(domain.LOG_CRIT, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}
	return true, logData
}

// LogError is a function that logs an error message.
//
// It takes in the priority of the error, a generic error message,
// and a specific error message as parameters.
// It returns a boolean indicating if the log was successfully saved,
// and the details of the log.
func LogError(priority domain.Priority, genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	logData := services.Details(domain.LOG_ERR, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}
	return true, logData
}

// LogAlert generates a log alert and saves it in the database.
//
// Parameters:
// - genericErrMsg: a string representing the generic error message.
// - errMsg: a string representing the specific error message.
//
// Returns:
// - bool: a boolean indicating whether the log alert was saved in the database.
// - LogDetails: a struct containing the details of the generated log alert.
func LogAlert(genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	logData := services.Details(domain.LOG_ALERT, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}
	return true, logData
}

// LogWarning is a function that logs a warning message and returns a boolean value and the log details.
//
// It takes two parameters: genericErrMsg (string) - the generic error message, and errMsg (string) - the specific error message.
// It returns a boolean value indicating whether the log was saved successfully and the log details (domain.LogDetails).
func LogWarning(genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	logData := services.Details(domain.LOG_WARNING, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}
	return true, logData
}

// LogNotice is a Go function that logs a notice message with error details.
//
// It takes two parameters:
// - genericErrMsg: a string that represents a generic error message.
// - errMsg: a string that represents the specific error message.
//
// It returns a boolean value indicating whether the log was saved successfully,
// and a domain.LogDetails struct that contains the details of the log.
func LogNotice(genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	logData := services.Details(domain.LOG_NOTICE, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}
	return true, logData
}

// LogInfo is a Go function that logs information.
//
// It takes in two parameters: genericErrMsg (string) and errMsg (string).
// It returns a boolean value and an instance of domain.LogDetails.
func LogInfo(genericErrMsg string, errMsg string) (bool, domain.LogDetails) {
	logData := services.Details(domain.LOG_INFO, genericErrMsg, errMsg)

	logDb := services.CheckEnvironment()
	if logDb {
		services.SaveLog(logData)
	}
	return true, logData
}

// LogDebug is a Go function that logs a debug message.
//
// It takes two parameters:
// - genericErrMsg: a string representing a generic error message
// - errMsg: a string representing the specific error message
//
// It returns a boolean value.
func LogDebug(genericErrMsg string, errMsg string) bool {
	services.Debug(domain.LOG_DEBUG, genericErrMsg, errMsg)
	return true
}
