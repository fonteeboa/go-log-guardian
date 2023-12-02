package domain

import "time"

// Priority represents the severity level for the logs
type Priority int

// Severity levels based on syslog.h
const (
	LOG_EMERG   Priority = iota // 0
	LOG_ALERT                   // 1
	LOG_CRIT                    // 2
	LOG_ERR                     // 3
	LOG_WARNING                 // 4
	LOG_NOTICE                  // 5
	LOG_INFO                    // 6
	LOG_DEBUG                   // 7
)

// Map to convert Priority to string
var PriorityToString = map[Priority]string{
	LOG_EMERG:   "LOG_EMERG",
	LOG_ALERT:   "LOG_ALERT",
	LOG_CRIT:    "LOG_CRIT",
	LOG_ERR:     "LOG_ERR",
	LOG_WARNING: "LOG_WARNING",
	LOG_NOTICE:  "LOG_NOTICE",
	LOG_INFO:    "LOG_INFO",
	LOG_DEBUG:   "LOG_DEBUG",
}

// LogBase represents the common fields for all types of logs
type LogBase struct {
	Priority            Priority  // Log priority
	LogLevel            string    // Log level (debug, info, error, etc.)
	Timestamp           time.Time // Time when the log was registered
	GenericErrorMessage string    // Generic log message
	ErrorMessage        string    // Specific error message
}

// LogFunction represents a log record related to a specific function
type LogFunction struct {
	LogBase             // Incorporates LogBase fields
	FunctionName string // Name of the registered function
}

// LogDatabase represents a log record related to database operations
type LogDatabase struct {
	LogBase          // Incorporates LogBase fields
	TableName string // Name of the table related to the operation
	Query     string // Query executed in the database
}

// LogRequest represents a log record related to requests
type LogRequest struct {
	LogBase             // Incorporates LogBase fields
	Method       string // HTTP method of the request (GET, POST, etc.)
	StatusCode   int    // HTTP status code
	Path         string // Path of the request
	ResponseSize int    // Response size
}
