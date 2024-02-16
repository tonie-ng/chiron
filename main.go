package main

import (
	"fmt"
	"os"
	"os/signal"
	"os/user"

	"chiron/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Printf("Error getting current user: %v\n", err)
		os.Exit(1)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		fmt.Println("\nReceived interrupt signal. Exiting...")
		os.Exit(0)
	}()

	args := os.Args

	repl.Load(args, user.Username)
}
