package webapp

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"path"

	"github.com/m-kostrzewa/powershell-for-programmers/core"
)

type WebApp struct {
	server *http.Server
	Mux    *http.ServeMux
}

type questionsListPage struct {
	Questions []core.Question
}

func NewWebApp(templatesDir string, questions []core.Question) *WebApp {
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
	for index, q := range questions {
		path := fmt.Sprintf("/questions/%v", index)
		questionToServe := q
		w.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			tmplQuestion.Execute(w, questionToServe)
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
