package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fonteeBoa/go-log-guardian/pkg"
	"github.com/fonteeBoa/go-log-guardian/pkg/domain"
)

func main() {
	dbConfigs := map[string]map[string]string{
		"postgres": {
			"POSTGRES_HOST":          "127.0.0.1",   // Safe: Localhost address for testing
			"POSTGRES_EXTERNAL_PORT": "5432",        // Safe: Standard PostgreSQL port
			"POSTGRES_USER":          "admin",       // Safe: Test user for local database
			"POSTGRES_PASSWORD":      "admin",       // Safe: Test password for local database
			"POSTGRES_DB":            "logguardian", // Safe: Test database name
		},
		"mysql": {
			"MYSQL_HOST":     "127.0.0.1",   // Safe: Localhost address for testing
			"MYSQL_PORT":     "3306",        // Safe: Standard MySQL port
			"MYSQL_USER":     "admin",       // Safe: Test user for local database
			"MYSQL_PASSWORD": "admin",       // Safe: Test password for local database
			"MYSQL_DBNAME":   "logguardian", // Safe: Test database name
		},
		"sqlite": {
			"SQLITE_PATH": "./data/logguardian.db", // Safe: Local SQLite file for testing
		},
		"mongodb": {
			"MONGODB_URI":    "mongodb://127.0.0.1:27017", // Safe: Local MongoDB URI for testing
			"MONGODB_DBNAME": "logguardian",               // Safe: Test database name
		},
		"elastic": {
			"ELASTIC_URI": "http://127.0.0.1:9200", // Safe: Local ElasticSearch URI for testing
		},
	}

	for dbType, config := range dbConfigs {
		fmt.Printf("\n=== Testing with database type: %s ===\n", dbType)

		setEnvVariables(config)

		if err := os.Setenv("DATABASE_TYPE", dbType); err != nil {
			log.Fatalf("Failed to set DATABASE_TYPE: %v", err)
		}

		runTestsForDatabase()
	}
}

func setEnvVariables(config map[string]string) {
	for key, value := range config {
		if err := os.Setenv(key, value); err != nil {
			log.Fatalf("Failed to set environment variable %s: %v", key, err)
		}
	}
}

func runTestsForDatabase() {
	fmt.Println("== Testing LogFunction ==")
	success, logFunction := pkg.LogFunction(
		domain.LOG_INFO,
		"TestFunction",
		"Generic function error",
		"Specific function error",
	)
	printResult(success, logFunction)

	fmt.Println("\n== Testing LogDataBase ==")
	success, logDatabase := pkg.LogDataBase(
		domain.LOG_ERR,
		"TestTable",
		"SELECT * FROM TestTable",
		"Generic database error",
		"Specific database error",
	)
	printResult(success, logDatabase)

	fmt.Println("\n== Testing LogRequests ==")
	success, logRequest := pkg.LogRequests(
		domain.LOG_WARNING,
		"GET",
		200,
		"/test-endpoint",
		1024,
		"Generic request error",
		"Specific request error",
	)
	printResult(success, logRequest)

	fmt.Println("\n== Testing Log ==")
	success, logDetails := pkg.Log(
		domain.LOG_ALERT,
		"Generic log error",
		"Specific log error",
	)
	printResult(success, logDetails)

	testSpecificLogFunctions()
}

func testSpecificLogFunctions() {
	logTests := []struct {
		name       string
		logFunc    func(string, string) (bool, domain.LogDetails)
		genericErr string
		errMsg     string
	}{
		{"LogEmerg", adaptLogFunction(pkg.LogEmerg), "Generic emerg error", "Specific emerg error"},
		{"LogCritical", adaptLogFunction(pkg.LogCritical), "Generic critical error", "Specific critical error"},
		{"LogAlert", adaptLogFunction(pkg.LogAlert), "Generic alert", "Specific alert"},
		{"LogWarning", adaptLogFunction(pkg.LogWarning), "Generic warning", "Specific warning"},
		{"LogNotice", adaptLogFunction(pkg.LogNotice), "Generic notice", "Specific notice"},
		{"LogInfo", adaptLogFunction(pkg.LogInfo), "Generic info", "Specific info"},
		{"LogError", func(genericErr, specificErr string) (bool, domain.LogDetails) {
			return pkg.LogError(domain.LOG_ERR, genericErr, specificErr)
		}, "Generic error", "Specific error"},
	}

	for _, test := range logTests {
		fmt.Printf("\n--- Testing %s ---\n", test.name)
		success, logDetails := test.logFunc(test.genericErr, test.errMsg)
		printResult(success, logDetails)
	}

	fmt.Println("\n--- Testing LogDebug ---")
	success := pkg.LogDebug("Generic debug error", "Specific debug error")
	fmt.Printf("LogDebug Success: %v\n", success)
}

func adaptLogFunction(f func(string, string) (bool, domain.LogDetails)) func(string, string) (bool, domain.LogDetails) {
	return func(genericErrMsg, errMsg string) (bool, domain.LogDetails) {
		return f(genericErrMsg, errMsg)
	}
}

func printResult(success bool, logData interface{}) {
	if !success {
		log.Println("Log operation failed!")
		return
	}

	fmt.Println("Log operation successful!")
	fmt.Printf("Log Data: %+v\n", logData)
}
