package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/m-kostrzewa/powershell-for-programmers/webapp"
)

func main() {
	webapp := webapp.WebApp{}
	webapp.Serve("templates")
	defer webapp.Shutdown()

	waitForSigInt()
}

func waitForSigInt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Graceful shutdown")
}
