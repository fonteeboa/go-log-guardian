package pkg

import (
	"GoLogGuardian/pkg/domain"
	"fmt"
	"os"
)

// checkEnvironment checks the environment and returns a boolean value.
//
// This function does not take any parameters.
// It returns a boolean value indicating whether the "LOG_GUARDIAN_DATABASE_TYPE" environment variable is set or not.
func checkEnvironment() bool {
	insertDB := os.Getenv("LOG_GUARDIAN_DATABASE_TYPE")
	return insertDB != ""
}

// debug prints the provided error message if the priority is set to LOG_DEBUG.
//
// Parameters:
//   - priority: the priority level of the error message (domain.Priority)
//   - genericErrMsg: the generic error message (string)
//   - errMsg: the specific error message (string)
func debug(priority domain.Priority, genericErrMsg string, errMsg string) {
	if priority == domain.LOG_DEBUG {
		fmt.Println(domain.PriorityToString[priority] + ": " + genericErrMsg + " " + errMsg)
	}
}
