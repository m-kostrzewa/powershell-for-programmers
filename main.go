package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/m-kostrzewa/powershell-for-programmers/adapters/webapp"
	"github.com/m-kostrzewa/powershell-for-programmers/core/question"
)

func main() {
	questions := []question.Question{
		*question.New(question.NextQuestionID(),
			"Lexical scope",
			"Does Powershell do X?",
			"Some pseudocode here....",
			[]question.Answer{
				{Text: "Answer 1", IsCorrect: true},
				{Text: "Answer 2", IsCorrect: false},
				{Text: "Answer 3", IsCorrect: false},
			},
		),
		*question.New(question.NextQuestionID(),
			"Scopes in closures",
			"What is the expected output?",
			"Some other pseudocode here....",
			[]question.Answer{
				{Text: "aab", IsCorrect: true},
				{Text: "abb", IsCorrect: false},
				{Text: "aba", IsCorrect: false},
			},
		),
	}
	app := webapp.NewWebApp("./templates", questions)
	app.Serve("templates")
	defer app.Shutdown()

	waitForSigInt()
}

func waitForSigInt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Graceful shutdown")
}
