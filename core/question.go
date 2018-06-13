package core

type Answer struct {
	Text      string
	IsCorrect bool
}

type Question struct {
	Title        string
	QuestionText string
	CodeBody     string
	Answers      []Answer
}

func (q *Question) IsCorrect(answerId int) bool {
	return q.Answers[answerId].IsCorrect
}
