package webapp

import (
	"fmt"
	"html/template"
	"net/http"
	"path"

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
