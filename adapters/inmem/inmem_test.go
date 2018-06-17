package inmem_test

import (
	"testing"

	"github.com/m-kostrzewa/powershell-for-programmers/adapters/inmem"
	"github.com/m-kostrzewa/powershell-for-programmers/core/domain/question"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	var repo question.Repository
	repo = inmem.New()
	assert.Empty(t, repo.FindAll())
}

func TestStore(t *testing.T) {
	var repo question.Repository
	repo = inmem.New()
	q := question.New(question.NextID(), "A", "B", "C", nil)
	err := repo.Store(q)
	assert.NoError(t, err)
	assert.NotEmpty(t, repo.FindAll())
}

func TestFind(t *testing.T) {
	var repo question.Repository
	repo = inmem.New()

	guid := question.NextID()
	q := question.New(guid, "A", "B", "C", nil)
	err := repo.Store(q)

	found, err := repo.Find(guid)
	assert.NoError(t, err)
	assert.Equal(t, found.QuestionID, q.QuestionID)
}

func TestFindDoesntExist(t *testing.T) {
	var repo question.Repository
	repo = inmem.New()

	guid := question.NextID()

	found, err := repo.Find(guid)
	assert.Error(t, err)
	assert.Nil(t, found)
}
