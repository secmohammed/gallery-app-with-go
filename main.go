package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
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

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	// http.HandleFunc("/", handleRequests)
	http.ListenAndServe(":3000", router)
}
