package main

import (
	"fmt"
	"net/http"
	"os"
)

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	port := os.Getenv("PORT")

	if len(port) < 1 {
		port = "8080"
	}

	http.HandleFunc("/", serveStatic)
	fmt.Println("Listening on port", port)
	http.ListenAndServe(":"+port, nil)
}
