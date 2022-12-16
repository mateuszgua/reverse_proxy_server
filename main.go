package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	originServerHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[origin server] received request at: %s\n", time.Now())
		_, _ = fmt.Print(r, "origin server response")
	})
	fmt.Print(originServerHandler)

	reverseProxy := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[reverse proxy server] received request at %s\n", time.Now())
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8090"
	}

	port := fmt.Sprintf(":%s", httpPort)
	log.Printf("Starting server on http://localhost%s", port)

	log.Fatal(http.ListenAndServe(port, reverseProxy))
}
