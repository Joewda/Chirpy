package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	mux := http.NewServeMux()

	my_server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Printf("Serving on port %s\n", port)
	log.Fatal(my_server.ListenAndServe())
}
