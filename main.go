package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/m-kostrzewa/powershell-for-programmers/adapters/inmem"
	"github.com/m-kostrzewa/powershell-for-programmers/core/application/webapp"
	"github.com/m-kostrzewa/powershell-for-programmers/core/question"
)

func main() {
	questions := inmem.New()
	questions.Store(question.New(question.NextID(),
		"Lexical scope",
		"Does Powershell do X?",
		`
Write-Host 'Hello, world!'
$a = Get-Credential`,
		[]question.Answer{
			{Text: "Answer 1", IsCorrect: true},
			{Text: "Answer 2", IsCorrect: false},
			{Text: "Answer 3", IsCorrect: false},
		},
	))
	questions.Store(question.New(question.NextID(),
		"Scopes in closures",
		"What is the expected output?",
		"Some other pseudocode here....",
		[]question.Answer{
			{Text: "aab", IsCorrect: true},
			{Text: "abb", IsCorrect: false},
			{Text: "aba", IsCorrect: false},
		},
	))
	app := webapp.NewWebApp(".", questions)
	app.Serve()
	defer app.Shutdown()

	waitForSigInt()
}

func waitForSigInt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Graceful shutdown")
}
