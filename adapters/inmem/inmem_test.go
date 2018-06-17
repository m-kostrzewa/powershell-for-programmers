package inmem_test

import (
	"testing"

	"github.com/m-kostrzewa/powershell-for-programmers/adapters/inmem"
	"github.com/m-kostrzewa/powershell-for-programmers/core/question"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	var repo question.QuestionRepository
	repo = inmem.New()
	assert.Empty(t, repo.FindAll())
}
