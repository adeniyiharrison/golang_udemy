package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/adeniyiharrison/golang_udemy/16_microservice_test/handlers"
	"golang.org/x/net/context"
)

func main() {
	stdOutLogger := log.New(os.Stdout, "product-api", log.LstdFlags)
	newHello := handlers.NewHello(stdOutLogger)
	newGoodBye := handlers.NewGoodBye(stdOutLogger)

	newServeMux := http.NewServeMux()
	newServeMux.Handle("/", newHello)
	newServeMux.Handle("/goodbye", newGoodBye)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      newServeMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			stdOutLogger.Fatalln(err)
		}
	}()

	// this will block until message is received off of channel
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)

	sig := <-sigchan
	stdOutLogger.Println("Recieved terminate, gracefully shutdown", sig)

	deadline, _ := context.WithTimeout(context.Background(), 10*time.Second)

	s.Shutdown(deadline)
}
