package main

import (
	"html/template"
	"net/http"

	"github.com/m-kostrzewa/powershell-for-programmers/core"
)

func main() {

	tmpl := template.Must(template.ParseFiles("templates/layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	http.ListenAndServe(":8080", nil)
}
