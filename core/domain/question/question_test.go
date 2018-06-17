package question_test

import (
	"testing"

	"github.com/satori/go.uuid"

	"github.com/m-kostrzewa/powershell-for-programmers/core/domain/question"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	id := question.NextID()
	q := question.New(id, "Question?", "Text", "Body",
		[]question.Answer{
			{Text: "Answer 1", IsCorrect: true},
		},
	)
	assert.False(t, uuid.Equal(uuid.Nil, uuid.UUID(q.QuestionID)))
}

func TestAnswerCorrect(t *testing.T) {

	q := question.Question{}
	q.Answers = append(q.Answers, question.Answer{IsCorrect: true})
	q.Answers = append(q.Answers, question.Answer{IsCorrect: false})
	q.Answers = append(q.Answers, question.Answer{IsCorrect: false})

	result := q.IsCorrect(1)
	assert.False(t, result)
}
