/**
 * @file domain.go
 * @brief This file contains the domain models and constants used for logging.
 */

package domain

import "time"

// Priority represents the severity level for the logs
/**
 * @enum Priority
 * @brief Severity levels based on syslog.h
 */
type Priority int

// Severity levels based on syslog.h
const (
	LOG_EMERG   Priority = iota ///< Emergency: system is unusable
	LOG_CRIT                    ///< Critical: critical conditions
	LOG_ERR                     ///< Error: error conditions
	LOG_ALERT                   ///< Alert: action must be taken immediately
	LOG_WARNING                 ///< Warning: warning conditions
	LOG_NOTICE                  ///< Notice: normal but significant condition
	LOG_INFO                    ///< Informational: informational messages
	LOG_DEBUG                   ///< Debug: debug-level messages
)

// Map to convert Priority to string
/**
 * @var PriorityToString
 * @brief Map to convert Priority to string representation.
 */
var PriorityToString = map[Priority]string{
	LOG_EMERG:   "LOG_EMERG",
	LOG_CRIT:    "LOG_CRIT",
	LOG_ERR:     "LOG_ERR",
	LOG_ALERT:   "LOG_ALERT",
	LOG_WARNING: "LOG_WARNING",
	LOG_NOTICE:  "LOG_NOTICE",
	LOG_INFO:    "LOG_INFO",
	LOG_DEBUG:   "LOG_DEBUG",
}

// LogDetails represents the common fields for all types of logs
/**
 * @struct LogDetails
 * @brief Represents the common fields for all types of logs.
 */
type LogDetails struct {
	Priority            Priority  ///< Log priority
	LogLevel            string    ///< Log level (debug, info, error, etc.)
	Timestamp           time.Time ///< Time when the log was registered
	GenericErrorMessage string    ///< Generic log message
	ErrorMessage        string    ///< Specific error message
}

// LogFunction represents a log record related to a specific function
/**
 * @struct LogFunction
 * @brief Represents a log record related to a specific function.
 * @extends LogDetails
 */
type LogFunction struct {
	LogDetails          ///< Incorporates LogDetails fields
	FunctionName string ///< Name of the registered function
}

// LogDatabase represents a log record related to database operations
/**
 * @struct LogDatabase
 * @brief Represents a log record related to database operations.
 * @extends LogDetails
 */
type LogDatabase struct {
	LogDetails        ///< Incorporates LogDetails fields
	TableName  string ///< Name of the table related to the operation
	Query      string ///< Query executed in the database
}

// LogRequest represents a log record related to requests
/**
 * @struct LogRequest
 * @brief Represents a log record related to requests.
 * @extends LogDetails
 */
type LogRequest struct {
	LogDetails          ///< Incorporates LogDetails fields
	Method       string ///< HTTP method of the request (GET, POST, etc.)
	StatusCode   int    ///< HTTP status code
	Path         string ///< Path of the request
	ResponseSize int    ///< Response size
}
