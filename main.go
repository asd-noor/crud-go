package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	api "github.com/darrSonik/crud-go/api"
	book "github.com/darrSonik/crud-go/models"
)

func main() {
	mockData := book.CreateMockData()
	serverFunc(mockData)
}

func serverFunc(data book.Books) {
	// fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc(
		"/books/",
		func(w http.ResponseWriter, r *http.Request) { api.HandlerBook(&data, w, r) },
	)

	http.HandleFunc(
		"/books",
		func(w http.ResponseWriter, r *http.Request) { api.HandlerBooks(&data, w, r) },
	)

	port := getPort()
	httpServer := http.Server{Addr: port}

	connectionClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := httpServer.Shutdown(context.Background()); err != nil {
			log.Printf("Server shutdown error: %+v", err)
		}
		close(connectionClosed)
	}()

	log.Printf("Listening on port %s", strings.TrimPrefix(port, ":"))
	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe Error: %+v", err)
	}
	<-connectionClosed

	log.Printf("Server closed\n\n")
}

func getPort() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter port number: ")
	port, _ := reader.ReadString('\n')

	return ":" + strings.TrimSuffix(port, "\n")
}
