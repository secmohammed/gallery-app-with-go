package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// global variable to have the template parsed only once, and execute it only when needed (much more performant.) than parsing each time.
// declaring multiple variables.
var (
	homeTemplate    *template.Template
	contactTemplate *template.Template
)

func handleHomeRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// panic only when executing the template has an error that's not null.
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}

}
func handleContactRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(response, nil); err != nil {
		return
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

	var err error
	// we don't use :=, to assign it to the the first global variable we created.
	homeTemplate, err = template.ParseFiles(
		"views/home.gohtml",
		"views/layouts/footer.gohtml",
	)
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles(
		"views/contact.gohtml",
		"views/layouts/footer.gohtml",
	)
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
