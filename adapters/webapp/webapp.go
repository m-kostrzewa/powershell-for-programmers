package webapp

import (
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

func NewWebApp(layoutsDir string) *WebApp {
	layoutTmpl := path.Join(layoutsDir, "layout.html")
	tmpl := template.Must(template.ParseFiles(layoutTmpl))

	w := WebApp{
		server: nil,
		Mux:    http.NewServeMux(),
	}

	w.Mux.HandleFunc("/questions", func(w http.ResponseWriter, r *http.Request) {
		data := core.Question{
			Title: "Lexical scope",
			Text:  "Does Powershell do X?",
			Body:  "Some pseudocode here....",
			Answers: []core.Answer{
				{Text: "Answer 1", IsCorrect: false},
				{Text: "Answer 2", IsCorrect: true},
				{Text: "Answer 3", IsCorrect: true},
			},
		}
		tmpl.Execute(w, data)
	})

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
