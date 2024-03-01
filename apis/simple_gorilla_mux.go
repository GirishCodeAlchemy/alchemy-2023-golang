package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

///handle example
type countHandler struct {
	mu sync.Mutex
	n int
}

// count the number of requests processed
func (h *countHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w,"\nTotal no of hit,%d\n",h.n)
}



func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to Home page </h1>")

}

func getData(w http.ResponseWriter, r *http.Request) {
	html := `<html><head></head><body>
	<form action='' method='POST'>
	The field <input name='field'/>
	</form></body></html>`
	w.Write([]byte(html))
}

func postData(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.FormValue("field")))
}

func pathMatcher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte(vars["id"]))
}

func queryMatcher(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL)
	// w.Write([]byte(r.URL))
}

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/", home)
	r.Handle("/count", new(countHandler))

	//Method matcher
	r.HandleFunc("/data", getData).Methods("GET")
	r.HandleFunc("/data", postData).Methods("POST")

	//pathMatcher
	r.HandleFunc("/data/{id:[0-9]+}", pathMatcher)

	//QUery matcher
	r.Path("/").Queries("id", "test1").HandlerFunc(queryMatcher)
	r.Path("/").Queries("id", "{id:[0-9]+}").HandlerFunc(queryMatcher)

	// Subroute
	sr := r.PathPrefix("/data1").Subrouter()
	sr.HandleFunc("/first", queryMatcher)

	fmt.Println("starting server")
	go http.ListenAndServe(":8000", r)
	fmt.Println("started server")
	fmt.Scanln()
}
