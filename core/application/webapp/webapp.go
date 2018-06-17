package webapp

import (
	"net"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/m-kostrzewa/powershell-for-programmers/core/domain/question"
)

type WebApp struct {
	server *http.Server
	Mux    *http.ServeMux
}

type answerForm struct {
	AnswerID int `json:"answerid"`
}

func NewWebApp(rootDir string, questionsRepo question.Repository) *WebApp {
	w := WebApp{
		server: nil,
		Mux:    http.NewServeMux(),
	}

	fs := http.FileServer(http.Dir(path.Join(rootDir, "static")))
	w.Mux.Handle("/static/", http.StripPrefix("/static/", fs))

	builder := NewBuilder(rootDir, "templates")
	qv := NewQuestionsController(builder, questionsRepo)
	gorillaMux := mux.NewRouter()
	gorillaMux.HandleFunc("/questions", qv.QuestionListHandler)
	gorillaMux.HandleFunc("/questions/{title}", qv.QuestionShowHandler)

	w.Mux.Handle("/", gorillaMux)

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
