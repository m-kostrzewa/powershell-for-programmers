package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/m-kostrzewa/powershell-for-programmers/adapters/webapp"
	"github.com/m-kostrzewa/powershell-for-programmers/core"
)

func main() {
	questions := []core.Question{
		{
			Title: "Lexical scope",
			Text:  "Does Powershell do X?",
			Body:  "Some pseudocode here....",
			Answers: []core.Answer{
				{Text: "Answer 1", IsCorrect: false},
				{Text: "Answer 2", IsCorrect: true},
				{Text: "Answer 3", IsCorrect: true},
			},
		},
		{
			Title: "Scopes in closures",
			Text:  "What is the expected output?",
			Body:  "Some other pseudocode here....",
			Answers: []core.Answer{
				{Text: "aab", IsCorrect: false},
				{Text: "abb", IsCorrect: true},
				{Text: "aba", IsCorrect: true},
			},
		},
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
