package webapp

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/gosimple/slug"
	"github.com/m-kostrzewa/powershell-for-programmers/core/domain/question"
)

type Builder struct {
	rootDir      string
	templatesDir string
}

func NewBuilder(rootDir, templatesDir string) *Builder {
	return &Builder{
		rootDir:      rootDir,
		templatesDir: templatesDir,
	}
}

func (v *Builder) Get(tmplFilename string) string {
	return path.Join(v.rootDir, v.templatesDir, tmplFilename)
}

type QuestionsController struct {
	builder       *Builder
	questionsRepo question.Repository
}

func NewQuestionsController(builder *Builder, questionsRepo question.Repository) *QuestionsController {
	return &QuestionsController{
		builder:       builder,
		questionsRepo: questionsRepo,
	}
}

func (qv *QuestionsController) QuestionListHandler(w http.ResponseWriter, r *http.Request) {
	layoutTmpl := qv.builder.Get("layout.html")
	qListTmpl := qv.builder.Get("questions_list.html")

	listQ := listQuestionsModel{
		Questions: []listQuestionsItemModel{},
	}
	questions := qv.questionsRepo.FindAll()

	for _, q := range questions {
		listQ.Questions = append(listQ.Questions,
			listQuestionsItemModel{
				Title: q.Title,
				Path:  "questions/" + slug.Make(q.Title),
			})
	}

	t := template.Must(template.ParseFiles(layoutTmpl, qListTmpl))
	err := t.ExecuteTemplate(w, "layout", listQ)
	if err != nil {
		fmt.Println(err)
	}
}

type listQuestionsItemModel struct {
	Title string
	Path  string
}

type listQuestionsModel struct {
	Questions []listQuestionsItemModel
}

func (qv *QuestionsController) QuestionShowHandler(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.EscapedPath(), "/")
	title := s[len(s)-1]

	layoutTmpl := qv.builder.Get("layout.html")
	qShowTmpl := qv.builder.Get("question.html")
	var subViewTmpl string

	showQ := showQuestionModel{
		Answers: []answerModel{},
	}
	// TODO: this is ugly
	questions := qv.questionsRepo.FindAll()

	for _, q := range questions {
		if slug.Make(q.Title) == title {
			showQ.Title = q.Title
			showQ.Text = q.Text
			showQ.Body = q.Body
			if r.Method == "POST" {
				r.ParseForm()
				answerID, _ := strconv.Atoi(r.Form.Get("answerID"))
				subViewTmpl = qv.builder.Get("congrats.html")
				if q.IsCorrect(answerID) {
					subViewTmpl = qv.builder.Get("congrats.html")
				} else {
					subViewTmpl = qv.builder.Get("condolences.html")
				}
			} else {
				for _, a := range q.Answers {
					subViewTmpl = qv.builder.Get("answerform.html")
					showQ.Answers = append(showQ.Answers, answerModel{Text: a.Text})
				}
			}
			break
		}
	}

	t := template.Must(template.ParseFiles(layoutTmpl, qShowTmpl, subViewTmpl))
	err := t.ExecuteTemplate(w, "layout", showQ)
	if err != nil {
		fmt.Println(err)
	}
}

type showQuestionModel struct {
	Title   string
	Text    string
	Body    string
	Answers []answerModel
}

type answerModel struct {
	Text string
}
