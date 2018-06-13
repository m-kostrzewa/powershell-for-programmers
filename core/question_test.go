package core_test

import (
	"testing"

	"github.com/m-kostrzewa/powershell-for-programmers/core"
	"github.com/stretchr/testify/assert"
)

func TestAnswerCorrect(t *testing.T) {

	q := core.Question{}
	q.Answers = append(q.Answers, core.Answer{IsCorrect: true})
	q.Answers = append(q.Answers, core.Answer{IsCorrect: false})
	q.Answers = append(q.Answers, core.Answer{IsCorrect: false})

	result := q.IsCorrect(1)
	assert.False(t, result)
}
