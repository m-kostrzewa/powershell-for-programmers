package webapp

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"path"
	"strconv"

	"github.com/m-kostrzewa/powershell-for-programmers/core/question"
)

type WebApp struct {
	server *http.Server
	Mux    *http.ServeMux
}

type answerForm struct {
	AnswerID int `json:"answerid"`
}

type questionsListPage struct {
	Questions []question.Question
}

func NewWebApp(templatesDir string, questions []question.Question) *WebApp {
	w := WebApp{
		server: nil,
		Mux:    http.NewServeMux(),
	}

	tmplQuestionsList := template.Must(template.ParseFiles(path.Join(templatesDir, "questions_list.html")))
	w.Mux.HandleFunc("/questions", func(w http.ResponseWriter, r *http.Request) {
		data := questionsListPage{
			Questions: questions,
		}
		tmplQuestionsList.Execute(w, data)
	})

	tmplQuestion := template.Must(template.ParseFiles(path.Join(templatesDir, "question.html")))
	tmplCongrats := template.Must(template.ParseFiles(path.Join(templatesDir, "congrats.html")))
	tmplCondolences := template.Must(template.ParseFiles(path.Join(templatesDir, "condolences.html")))
	for index, q := range questions {
		questionToServe := q

		questionPath := fmt.Sprintf("/questions/%v", index)
		w.Mux.HandleFunc(questionPath, func(w http.ResponseWriter, r *http.Request) {

			tmplQuestion.Execute(w, questionToServe)
			if r.Method == "POST" {
				r.ParseForm()
				answerID, _ := strconv.Atoi(r.Form.Get("answerID"))
				if questionToServe.IsCorrect(answerID) {
					tmplCongrats.Execute(w, nil)
				} else {
					tmplCondolences.Execute(w, nil)
				}
			}
		})
	}

	return &w
}

func (w *WebApp) Serve(layoutsDir string) {
	listener, _ := net.ListenTCP("tcp4", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080})
	w.server = &http.Server{Handler: w.Mux}

	go func() {
		_ = w.server.Serve(listener)
	}()
}

func (w *WebApp) Shutdown() {
	w.server.Shutdown(nil)
}
