package pkg_test

import (
	"testing"

	"github.com/fonteeBoa/go-log-guardian/internal/testUtils/mocks"
	"github.com/fonteeBoa/go-log-guardian/pkg"
	"github.com/fonteeBoa/go-log-guardian/pkg/domain"
)

func TestLogFunction(t *testing.T) {
	expectedLog := domain.LogFunction{
		LogDetails: domain.LogDetails{
			Priority:            domain.LOG_ERR,
			LogLevel:            "debug",
			GenericErrorMessage: mocks.MockGeneralError,
			ErrorMessage:        mocks.MockSpecificError,
		},
		FunctionName: "testFunction",
	}

	success, result := pkg.LogFunction(domain.LOG_ERR, "testFunction", mocks.MockGeneralError, mocks.MockSpecificError)

	if !success {
		t.Errorf("LogFunction failed to save log")
	}

	if result.Priority != expectedLog.Priority ||
		result.GenericErrorMessage != expectedLog.GenericErrorMessage ||
		result.ErrorMessage != expectedLog.ErrorMessage ||
		result.FunctionName != expectedLog.FunctionName {
		t.Errorf("LogFunction did not return the expected log")
		t.Errorf("Expected: %v", expectedLog)
		t.Errorf("Result: %v", result)
	}
}

func TestLogDataBase(t *testing.T) {
	expectedLog := domain.LogDatabase{
		LogDetails: domain.LogDetails{
			Priority:            domain.LOG_ERR,
			LogLevel:            "debug",
			GenericErrorMessage: mocks.MockGeneralError,
			ErrorMessage:        mocks.MockSpecificError,
		},
		TableName: mocks.MockTable,
		Query:     "testQuery",
	}

	success, result := pkg.LogDataBase(domain.LOG_ERR, mocks.MockTable, "testQuery", mocks.MockGeneralError, mocks.MockSpecificError)

	if !success {
		t.Errorf("LogDataBase failed to save log")
	}

	if result.Priority != expectedLog.Priority ||
		result.GenericErrorMessage != expectedLog.GenericErrorMessage ||
		result.ErrorMessage != expectedLog.ErrorMessage ||
		result.TableName != expectedLog.TableName ||
		result.Query != expectedLog.Query {
		t.Errorf("LogDataBase did not return the expected log")
		t.Errorf("Expected: %v", expectedLog)
		t.Errorf("Result: %v", result)
	}
}

func TestLogRequests(t *testing.T) {
	expectedLog := domain.LogRequest{
		LogDetails: domain.LogDetails{
			Priority:            domain.LOG_ERR,
			LogLevel:            "debug",
			GenericErrorMessage: mocks.MockGeneralError,
			ErrorMessage:        mocks.MockSpecificError,
		},
		Method:       "GET",
		StatusCode:   200,
		Path:         "/test",
		ResponseSize: 100,
	}

	success, result := pkg.LogRequests(domain.LOG_ERR, "GET", 200, "/test", 100, mocks.MockGeneralError, mocks.MockSpecificError)

	if !success {
		t.Errorf("LogRequests failed to save log")
	}

	if result.Priority != expectedLog.Priority ||
		result.GenericErrorMessage != expectedLog.GenericErrorMessage ||
		result.ErrorMessage != expectedLog.ErrorMessage ||
		result.Method != expectedLog.Method ||
		result.StatusCode != expectedLog.StatusCode ||
		result.Path != expectedLog.Path ||
		result.ResponseSize != expectedLog.ResponseSize {
		t.Errorf("LogRequests did not return the expected log")
	}
}

// Teste para a função Log
func TestLog(t *testing.T) {
	expectedLog := domain.LogDetails{
		Priority:            domain.LOG_DEBUG,
		LogLevel:            "debug",
		GenericErrorMessage: mocks.MockGeneralError,
		ErrorMessage:        mocks.MockSpecificError,
	}

	success, result := pkg.Log(domain.LOG_DEBUG, mocks.MockGeneralError, mocks.MockSpecificError)

	if !success {
		t.Errorf("Log failed to save log")
	}

	if result.Priority != expectedLog.Priority ||
		result.GenericErrorMessage != expectedLog.GenericErrorMessage ||
		result.ErrorMessage != expectedLog.ErrorMessage {
		t.Errorf("Log did not return the expected log")
	}
}

// Teste para a função LogEmerg
func TestLogEmerg(t *testing.T) {
	expectedLog := domain.LogDetails{
		Priority:            domain.LOG_EMERG,
		LogLevel:            "LOG_EMERG",
		GenericErrorMessage: mocks.MockGeneralError,
		ErrorMessage:        mocks.MockSpecificError,
	}

	success, result := pkg.LogEmerg(mocks.MockGeneralError, mocks.MockSpecificError)

	if !success {
		t.Errorf("LogEmerg failed to save log")
	}

	if result.Priority != expectedLog.Priority ||
		result.GenericErrorMessage != expectedLog.GenericErrorMessage ||
		result.ErrorMessage != expectedLog.ErrorMessage {
		t.Errorf("LogEmerg did not return the expected log")
	}
}

// Teste para a função LogCritical
func TestLogCritical(t *testing.T) {
	expectedLog := domain.LogDetails{
		Priority:            domain.LOG_CRIT,
		LogLevel:            "LOG_CRIT",
		GenericErrorMessage: mocks.MockGeneralError,
		ErrorMessage:        mocks.MockSpecificError,
	}

	success, result := pkg.LogCritical(mocks.MockGeneralError, mocks.MockSpecificError)

	if !success {
		t.Errorf("LogCritical failed to save log")
	}

	if result.Priority != expectedLog.Priority ||
		result.GenericErrorMessage != expectedLog.GenericErrorMessage ||
		result.ErrorMessage != expectedLog.ErrorMessage {
		t.Errorf("LogCritical did not return the expected log")
	}
}

// Teste para a função LogError
func TestLogError(t *testing.T) {
	expectedLog := domain.LogDetails{
		Priority:            domain.LOG_ERR,
		LogLevel:            "LOG_ERR",
		GenericErrorMessage: mocks.MockGeneralError,
		ErrorMessage:        mocks.MockSpecificError,
	}

	success, result := pkg.LogError(domain.LOG_ERR, mocks.MockGeneralError, mocks.MockSpecificError)

	if !success {
		t.Errorf("LogError failed to save log")
	}

	if result.Priority != expectedLog.Priority ||
		result.GenericErrorMessage != expectedLog.GenericErrorMessage ||
		result.ErrorMessage != expectedLog.ErrorMessage {
		t.Errorf("LogError did not return the expected log")
	}
}

// Teste para a função LogAlert
func TestLogAlert(t *testing.T) {
	expectedLog := domain.LogDetails{
		Priority:            domain.LOG_ALERT,
		LogLevel:            "LOG_ALERT",
		GenericErrorMessage: mocks.MockGeneralError,
		ErrorMessage:        mocks.MockSpecificError,
	}

	success, result := pkg.LogAlert(mocks.MockGeneralError, mocks.MockSpecificError)

	if !success {
		t.Errorf("LogAlert failed to save log")
	}

	if result.Priority != expectedLog.Priority ||
		result.GenericErrorMessage != expectedLog.GenericErrorMessage ||
		result.ErrorMessage != expectedLog.ErrorMessage {
		t.Errorf("LogAlert did not return the expected log")
	}
}

// Teste para a função LogWarning
func TestLogWarning(t *testing.T) {
	expectedLog := domain.LogDetails{
		Priority:            domain.LOG_WARNING,
		LogLevel:            "LOG_WARNING",
		GenericErrorMessage: mocks.MockGeneralError,
		ErrorMessage:        mocks.MockSpecificError,
	}

	success, result := pkg.LogWarning(mocks.MockGeneralError, mocks.MockSpecificError)

	if !success {
		t.Errorf("LogWarning failed to save log")
	}

	if result.Priority != expectedLog.Priority ||
		result.GenericErrorMessage != expectedLog.GenericErrorMessage ||
		result.ErrorMessage != expectedLog.ErrorMessage {
		t.Errorf("LogWarning did not return the expected log")
	}
}

// Teste para a função LogNotice
func TestLogNotice(t *testing.T) {
	expectedLog := domain.LogDetails{
		Priority:            domain.LOG_NOTICE,
		LogLevel:            "LOG_NOTICE",
		GenericErrorMessage: mocks.MockGeneralError,
		ErrorMessage:        mocks.MockSpecificError,
	}

	success, result := pkg.LogNotice(mocks.MockGeneralError, mocks.MockSpecificError)

	if !success {
		t.Errorf("LogNotice failed to save log")
	}

	if result.Priority != expectedLog.Priority ||
		result.GenericErrorMessage != expectedLog.GenericErrorMessage ||
		result.ErrorMessage != expectedLog.ErrorMessage {
		t.Errorf("LogNotice did not return the expected log")
	}
}

// Teste para a função LogInfo
func TestLogInfo(t *testing.T) {
	expectedLog := domain.LogDetails{
		Priority:            domain.LOG_INFO,
		LogLevel:            "LOG_INFO",
		GenericErrorMessage: mocks.MockGeneralError,
		ErrorMessage:        mocks.MockSpecificError,
	}

	success, result := pkg.LogInfo(mocks.MockGeneralError, mocks.MockSpecificError)

	if !success {
		t.Errorf("LogInfo failed to save log")
	}

	if result.Priority != expectedLog.Priority ||
		result.GenericErrorMessage != expectedLog.GenericErrorMessage ||
		result.ErrorMessage != expectedLog.ErrorMessage {
		t.Errorf("LogInfo did not return the expected log")
	}
}

// Teste para a função LogDebug
func TestLogDebug(t *testing.T) {
	success := pkg.LogDebug(mocks.MockGeneralError, mocks.MockSpecificError)

	if !success {
		t.Errorf("LogDebug failed")
	}
}
