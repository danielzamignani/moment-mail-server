package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", ExampleHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server is running on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error: could not start server: %v", err)
	}
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test")
}
