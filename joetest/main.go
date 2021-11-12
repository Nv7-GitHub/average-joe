package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	joe "github.com/Nv7-Github/average-joe"
)

//go:embed token.txt
var token string

func main() {
	joe, err := joe.NewJoe(token)
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal)
	done := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		fmt.Println("Listening, press CTRL+C to stop.")
		<-c
		fmt.Println("Cleaning up...")
		joe.Close()
		done <- true
	}()

	<-done
}
