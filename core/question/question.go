package question

import (
	"github.com/satori/go.uuid"
)

type Answer struct {
	Text      string
	IsCorrect bool
}

type QuestionID uuid.UUID

type Question struct {
	QuestionID QuestionID
	Title      string
	Text       string
	Body       string
	Answers    []Answer
}

type QuestionRepository interface {
	Find(QuestionID) (*Question, error)
	FindAll() []*Question
	Store(*Question) error
}

func New(questionID QuestionID, title, text, body string, answers []Answer) *Question {
	return &Question{
		QuestionID: questionID,
		Title:      title,
		Text:       text,
		Body:       body,
		Answers:    answers,
	}
}

func NextQuestionID() QuestionID {
	return QuestionID(uuid.Must(uuid.NewV4()))
}

func (q *Question) IsCorrect(answerId int) bool {
	return q.Answers[answerId].IsCorrect
}
