package services_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fonteeBoa/go-log-guardian/internal/services"
	"github.com/fonteeBoa/go-log-guardian/pkg/domain"
)

func TestCheckEnvironmentWithEnvVariableSet(t *testing.T) {
	// Set the environment variable for testing
	os.Setenv("LOG_GUARDIAN_DATABASE_TYPE", "test")

	envCheck := services.CheckEnvironment()

	assert.True(t, envCheck)
}

func TestCheckEnvironmentWithEnvVariableNotSet(t *testing.T) {
	// Unset the environment variable for testing
	os.Unsetenv("LOG_GUARDIAN_DATABASE_TYPE")

	envCheck := services.CheckEnvironment()

	assert.False(t, envCheck)
}

func TestDebugPrintsMessageForLogDebugPriority(t *testing.T) {
	// Redirect stdout to capture printed output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		// Restore stdout
		os.Stdout = old
	}()

	priority := domain.LOG_DEBUG
	genericErrMsg := "Test generic error"
	errMsg := "Test specific error"

	services.Debug(priority, genericErrMsg, errMsg)

	// Close the writer before reading the content
	w.Close()

	// Read the content from the pipe
	output := make([]byte, 1024)
	n, _ := r.Read(output)

	assert.Contains(t, string(output[:n]), "LOG_DEBUG: Test generic error Test specific error")
}

func TestDebugDoesNotPrintMessageForNonLogDebugPriority(t *testing.T) {
	// Redirect stdout to capture printed output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		// Restore stdout
		os.Stdout = old
	}()

	priority := domain.LOG_INFO
	genericErrMsg := "Test generic error"
	errMsg := "Test specific error"

	services.Debug(priority, genericErrMsg, errMsg)

	// Close the writer before reading the content
	w.Close()

	// Read the content from the pipe
	output := make([]byte, 1024)
	n, _ := r.Read(output)

	assert.Empty(t, string(output[:n]))
}
