package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/views"
)

// global variable to have the template parsed only once, and execute it only when needed (much more performant.) than parsing each time.
// declaring multiple variables.
var (
	homeView    *views.View
	contactView *views.View
)

func handleHomeRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	// panic only when executing the template has an error that's not null.
	if err != nil {
		panic(err)
	}

}
func handleContactRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(response, contactView.Layout, nil)

	if err != nil {
		panic(err)
	}

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
	homeView = views.NewView(
		"layout",
		"views/home.gohtml",
	)
	contactView = views.NewView(
		"layout",
		"views/contact.gohtml",
	)
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(customNotFoundPage)
	router.HandleFunc("/", handleHomeRequest)
	router.HandleFunc("/faq", handleFAQRequest)
	router.HandleFunc("/contact", handleContactRequest)
	http.ListenAndServe(":3000", router)
}
