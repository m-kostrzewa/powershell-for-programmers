package executor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/m-kostrzewa/powershell-for-programmers/adapters/executor"
)

func TestNew(t *testing.T) {
	executor, err := executor.New()
	assert.NoError(t, err)
	assert.NotNil(t, executor)
}

func TestExecute(t *testing.T) {
	executor, err := executor.New()
	out, err := executor.Execute("Write-Host 'hi'")
	assert.NoError(t, err)
	assert.Equal(t, "hi", out)
}
