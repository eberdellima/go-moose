package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-moose/database"
	"go-moose/src"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	db := database.InitializeDBConnection()
	defer db.Close()

	router := src.ConfigureAPI()

	errorChannel := make(chan error, 2)

	go listenToPort(errorChannel, router)
	go listenToSysCall(errorChannel)

	fmt.Printf("\nTerminated: %s\n", <-errorChannel)
}

func setupPort() string {
	port := "8000"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	return fmt.Sprintf(":%s", port)
}

func listenToPort(errorChannel chan<- error, handler http.Handler) {

	server := &http.Server{
		Addr:           setupPort(),
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	errorChannel <- server.ListenAndServe()
}

func listenToSysCall(errorChannel chan<- error) {
	systemChannel := make(chan os.Signal, 1)
	signal.Notify(systemChannel, syscall.SIGINT)
	errorChannel <- fmt.Errorf("%s", <-systemChannel)
}
