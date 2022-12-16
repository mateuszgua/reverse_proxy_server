package origin

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func OriginServer() {
	originServerHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[origin server] received request at: %s\n", time.Now())
		_, _ = fmt.Print(r, "origin server response")
	})
	fmt.Print(originServerHandler)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8091"
	}

	port := fmt.Sprintf(":%s", httpPort)
	log.Printf("Starting server on http://localhost%s", port)

	log.Fatal(http.ListenAndServe(port, originServerHandler))
}
