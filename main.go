package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	originServerURL, err := url.Parse("http://127.0.0.1:8081")
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	reverseProxy := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[reverse proxy server] received request at %s\n", time.Now())

		r.Host = originServerURL.Host
		r.URL.Host = originServerURL.Host
		r.URL.Scheme = originServerURL.Scheme
		r.RequestURI = ""

		originServerResponse, err := http.DefaultClient.Do(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		io.Copy(w, originServerResponse.Body)
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8090"
	}

	port := fmt.Sprintf(":%s", httpPort)
	log.Printf("Starting server on http://localhost%s", port)

	log.Fatal(http.ListenAndServe(port, reverseProxy))

}
