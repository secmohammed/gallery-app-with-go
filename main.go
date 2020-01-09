package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "Welcome to my awesome site.")

	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "welcome to the contact page.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Page 404 couldn't be found.")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleRequests)
	// http.HandleFunc("/", handleRequests)
	http.ListenAndServe(":3000", router)
}
