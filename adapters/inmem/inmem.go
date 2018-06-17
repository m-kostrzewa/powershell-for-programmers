package inmem

import "github.com/m-kostrzewa/powershell-for-programmers/core/question"

type InMemQuestionRepository struct {
	questions map[question.QuestionID]*question.Question
}

func New() *InMemQuestionRepository {
	return &InMemQuestionRepository{
		questions: map[question.QuestionID]*question.Question{},
	}
}

func (r *InMemQuestionRepository) Find(guid question.QuestionID) (*question.Question, error) {
	if val, ok := r.questions[guid]; ok {
		return val, nil
	} else {
		return nil, question.ErrNotFound
	}
}

func (r *InMemQuestionRepository) FindAll() []*question.Question {
	list := make([]*question.Question, 0)
	for _, val := range r.questions {
		list = append(list, val)
	}
	return list
}

func (r *InMemQuestionRepository) Store(q *question.Question) error {
	r.questions[q.QuestionID] = q
	return nil
}
