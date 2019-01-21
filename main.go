package main

import (
	"os"
	"os/signal"
	"syscall"
	"log"
	"github.com/breathman/golang-app-runners-example/run"
)

func main() {
	// get your app
	app, err := run.New()
	if err != nil {
		log.Fatal(err)
	}

	// run all functions from runners
	app.Run()

	ch := make(chan os.Signal, 10)
	signal.Notify(ch,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	defer close(ch)

	// wait when all will started
	<-app.Started

	// listen system channel for closing app
	go func() {
		sig, ok := <-ch
		if ok {
			log.Printf("shutdown process on %s system signal\n", sig)
		}
		app.Shutdown()
	}()

	// exist with correct status
	// if all slams with out error code code 0, in otherwise will 1
	os.Exit(<-app.Done)
}