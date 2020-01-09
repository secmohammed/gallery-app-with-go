package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1> Hello, we are at homepage </h1>")

}
func handleContactRequest(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Welcome to the contact page.")
}
func handleFAQRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "This is the FAQ Page.")
}
func customNotFoundPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Sorry, but we couldn't find the page you are looking for.")
}
func main() {
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(customNotFoundPage)
	router.HandleFunc("/", home)
	router.HandleFunc("/faq", handleFAQRequest)
	router.HandleFunc("/contact", handleContactRequest)
	http.ListenAndServe(":3000", router)
}
