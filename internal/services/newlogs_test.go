package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fonteeBoa/go-log-guardian/internal/services"
	"github.com/fonteeBoa/go-log-guardian/internal/testUtils/mocks"

	"github.com/fonteeBoa/go-log-guardian/pkg/domain"
)

func TestNewLogDetails(t *testing.T) {
	priority := domain.LOG_INFO
	genericErrMsg := mocks.MockGeneralError
	errMsg := mocks.MockSpecificError

	logDetails := services.NewLogDetails(priority, genericErrMsg, errMsg)

	assert.Equal(t, priority, logDetails.Priority)
	assert.Equal(t, "LOG_INFO", logDetails.LogLevel) // Verifica se a string do LogLevel está correta
	assert.NotZero(t, logDetails.Timestamp)          // Verifica se o Timestamp não é zero
	assert.Equal(t, genericErrMsg, logDetails.GenericErrorMessage)
	assert.Equal(t, errMsg, logDetails.ErrorMessage)
}

func TestNewFunctionLog(t *testing.T) {
	priority := domain.LOG_DEBUG
	functionName := "TestFunction"
	genericErrMsg := mocks.MockGeneralError
	errMsg := mocks.MockSpecificError

	functionLog := services.NewFunctionLog(priority, functionName, genericErrMsg, errMsg)

	assert.Equal(t, priority, functionLog.Priority)
	assert.Equal(t, "LOG_DEBUG", functionLog.LogLevel) // Verifica se a string do LogLevel está correta
	assert.NotZero(t, functionLog.Timestamp)           // Verifica se o Timestamp não é zero
	assert.Equal(t, genericErrMsg, functionLog.GenericErrorMessage)
	assert.Equal(t, errMsg, functionLog.ErrorMessage)
	assert.Equal(t, functionName, functionLog.FunctionName)
}

func TestNewDatabaseLog(t *testing.T) {
	priority := domain.LOG_NOTICE
	tableName := mocks.MockTable
	query := "SELECT * FROM TestTable"
	genericErrMsg := mocks.MockGeneralError
	errMsg := mocks.MockSpecificError

	databaseLog := services.NewDatabaseLog(priority, tableName, query, genericErrMsg, errMsg)

	assert.Equal(t, priority, databaseLog.Priority)
	assert.Equal(t, "LOG_NOTICE", databaseLog.LogLevel) // Verifica se a string do LogLevel está correta
	assert.NotZero(t, databaseLog.Timestamp)            // Verifica se o Timestamp não é zero
	assert.Equal(t, genericErrMsg, databaseLog.GenericErrorMessage)
	assert.Equal(t, errMsg, databaseLog.ErrorMessage)
	assert.Equal(t, tableName, databaseLog.TableName)
	assert.Equal(t, query, databaseLog.Query)
}

func TestNewRequestLog(t *testing.T) {
	priority := domain.LOG_CRIT
	method := "POST"
	statusCode := 500
	path := "/test/path"
	responseSize := 1024
	genericErrMsg := mocks.MockGeneralError
	errMsg := mocks.MockSpecificError

	requestLog := services.NewRequestLog(priority, method, statusCode, path, responseSize, genericErrMsg, errMsg)

	assert.Equal(t, priority, requestLog.Priority)
	assert.Equal(t, "LOG_CRIT", requestLog.LogLevel) // Verifica se a string do LogLevel está correta
	assert.NotZero(t, requestLog.Timestamp)          // Verifica se o Timestamp não é zero
	assert.Equal(t, genericErrMsg, requestLog.GenericErrorMessage)
	assert.Equal(t, errMsg, requestLog.ErrorMessage)
	assert.Equal(t, method, requestLog.Method)
	assert.Equal(t, statusCode, requestLog.StatusCode)
	assert.Equal(t, path, requestLog.Path)
	assert.Equal(t, responseSize, requestLog.ResponseSize)
}
