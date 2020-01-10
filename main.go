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
	signupView  *views.View
)

func handleHomeRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}
func handleContactRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	must(contactView.Render(response, nil))
}
func handleSignupRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	must(signupView.Render(response, nil))
}
func handleRegisterFormSubmission(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.PostFormValue("email"))
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
func must(err error) {
	if err != nil {
		panic(err)
	}
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
	signupView = views.NewView("layout", "views/signup.gohtml")
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(customNotFoundPage)
	router.HandleFunc("/", handleHomeRequest)
	router.HandleFunc("/register", handleRegisterFormSubmission).Methods("POST")
	router.HandleFunc("/faq", handleFAQRequest)
	router.HandleFunc("/contact", handleContactRequest)
	router.HandleFunc("/register", handleSignupRequest)
	http.ListenAndServe(":3000", router)
}
