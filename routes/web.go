package routes

import (
    "net/http"

    "github.com/gorilla/mux"
    "lenslocked.com/controllers"
)

// RegisterRoutes is used to register the routes we need for the web application.
func RegisterRoutes() {
    router := mux.NewRouter()
    router.NotFoundHandler = http.HandlerFunc(controllers.Show404Page().Render)
    router.HandleFunc("/", controllers.ShowHomePage().Render)
    router.HandleFunc("/contact", controllers.ShowContactPage().Render)
    // auth routes.
    auth := router.PathPrefix("/auth").Subrouter()
    auth.HandleFunc("/register", controllers.ParseRegisterForm).Methods("POST")
    auth.HandleFunc("/login", controllers.ParseLoginForm).Methods("POST")
    auth.HandleFunc("/login", controllers.ShowLoginForm().Render).Methods("GET")
    http.ListenAndServe(":3000", router)

}
