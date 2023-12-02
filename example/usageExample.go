package example

import (
	"fmt"

	"github.com/fonteeBoa/go-log-guardian/pkg"
	"github.com/fonteeBoa/go-log-guardian/pkg/domain"
)

func ShowExample() {
	// Example usage of LogFunction
	functionLogged, logFunctionData := pkg.LogFunction(domain.LOG_INFO, "ExampleFunction", "common.error", "Specific error message")
	fmt.Printf("Function logged: %v\nLogged data: %+v\n\n", functionLogged, logFunctionData)

	// Example usage of LogDataBase
	databaseLogged, logDatabaseData := pkg.LogDataBase(domain.LOG_ERR, "ExampleTable", "SELECT * FROM example_table", "common.error.get.example", "Specific error message")
	fmt.Printf("Database operation logged: %v\nLogged data: %+v\n\n", databaseLogged, logDatabaseData)

	// Example usage of LogRequests
	requestLogged, logRequestData := pkg.LogRequests(domain.LOG_NOTICE, "GET", 200, "/example", 1024, "common.error", "Specific error message")
	fmt.Printf("Request logged: %v\nLogged data: %+v\n\n", requestLogged, logRequestData)

	// Example usage of Log
	logged, logDetails := pkg.Log(domain.LOG_DEBUG, "common.error", "Specific error message")
	fmt.Printf("Log details: %v\nLogged data: %+v\n\n", logged, logDetails)

	// Example usage of LogEmerg
	emergencyLogged, logEmergencyDetails := pkg.LogEmerg("common.log.emergency", "Specific error message")
	fmt.Printf("Emergency log details: %v\nLogged data: %+v\n\n", emergencyLogged, logEmergencyDetails)

	// Example usage of LogCritical
	criticalLogged, logCriticalDetails := pkg.LogCritical("common.log.critical", "Specific error message")
	fmt.Printf("Critical log details: %v\nLogged data: %+v\n\n", criticalLogged, logCriticalDetails)

	// Example usage of LogError
	errorLogged, logErrorDetails := pkg.LogError(domain.LOG_ERR, "common.log.error", "Specific error message")
	fmt.Printf("Error log details: %v\nLogged data: %+v\n\n", errorLogged, logErrorDetails)

	// Example usage of LogAlert
	alertLogged, logAlertDetails := pkg.LogAlert("common.log.alert", "Specific error message")
	fmt.Printf("Alert log details: %v\nLogged data: %+v\n\n", alertLogged, logAlertDetails)

	// Example usage of LogWarning
	warningLogged, logWarningDetails := pkg.LogWarning("common.log.warning", "Specific error message")
	fmt.Printf("Warning log details: %v\nLogged data: %+v\n\n", warningLogged, logWarningDetails)

	// Example usage of LogNotice
	noticeLogged, logNoticeDetails := pkg.LogNotice("common.log.notice", "Specific error message")
	fmt.Printf("Notice log details: %v\nLogged data: %+v\n\n", noticeLogged, logNoticeDetails)

	// Example usage of LogInfo
	infoLogged, logInfoDetails := pkg.LogInfo("common.log.info", "Specific error message")
	fmt.Printf("Info log details: %v\nLogged data: %+v\n\n", infoLogged, logInfoDetails)

	// Example usage of LogDebug
	debugLogged := pkg.LogDebug("common.error", "Specific error message")
	fmt.Printf("Debug log: %v\n\n", debugLogged)
}
