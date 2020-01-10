package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template

func handleHomeRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}

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
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(customNotFoundPage)
	router.HandleFunc("/", handleHomeRequest)
	router.HandleFunc("/faq", handleFAQRequest)
	router.HandleFunc("/contact", handleContactRequest)
	http.ListenAndServe(":3000", router)
}
