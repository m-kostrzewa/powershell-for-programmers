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

func NewWebApp(rootDir string, questions []question.Question) *WebApp {
	w := WebApp{
		server: nil,
		Mux:    http.NewServeMux(),
	}

	fs := http.FileServer(http.Dir(path.Join(rootDir, "static")))
	w.Mux.Handle("/static/", http.StripPrefix("/static/", fs))

	w.Mux.HandleFunc("/questions", func(w http.ResponseWriter, r *http.Request) {
		layout := path.Join(rootDir, "templates", "layout.html")
		questionsLis := path.Join(rootDir, "templates", "questions_list.html")

		var tmpl *template.Template

		data := questionsListPage{
			Questions: questions,
		}

		tmpl = template.Must(template.ParseFiles(layout, questionsLis))
		tmpl.ExecuteTemplate(w, "layout", data)
	})

	for index, q := range questions {
		questionToServe := q

		questionPath := fmt.Sprintf("/questions/%v", index)
		w.Mux.HandleFunc(questionPath, func(w http.ResponseWriter, r *http.Request) {

			layout := path.Join(rootDir, "templates", "layout.html")
			question := path.Join(rootDir, "templates", "question.html")
			var tmpl *template.Template

			if r.Method == "POST" {
				r.ParseForm()
				answerID, _ := strconv.Atoi(r.Form.Get("answerID"))
				var result string
				if questionToServe.IsCorrect(answerID) {
					result = path.Join(rootDir, "templates", "congrats.html")
				} else {
					result = path.Join(rootDir, "templates", "condolences.html")
				}
				tmpl = template.Must(template.ParseFiles(layout, question, result))
			} else {
				tmpl = template.Must(template.ParseFiles(layout, question))
			}

			tmpl.ExecuteTemplate(w, "layout", questionToServe)
		})
	}

	return &w
}

func (w *WebApp) Serve() {
	listener, _ := net.ListenTCP("tcp4", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080})
	w.server = &http.Server{Handler: w.Mux}

	go func() {
		_ = w.server.Serve(listener)
	}()
}

func (w *WebApp) Shutdown() {
	w.server.Shutdown(nil)
}
