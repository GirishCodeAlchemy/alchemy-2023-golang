package lib

import (
	"fmt"
	"io"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		io.WriteString(w, "Hello, world! (GET)")
	case http.MethodPost:
		io.WriteString(w, "Hello, world! (POST)")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
	}
}

func SimpleWebserver() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world")
	})
	// http.HandleFunc("/", handleRoot)

	http.ListenAndServe(":8080", nil)
}
