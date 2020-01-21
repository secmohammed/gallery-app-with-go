package routes

import (
    "fmt"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "lenslocked.com/controllers"
)

// RegisterRoutes is used to register the routes we need for the web application.
func RegisterRoutes() {
    router := mux.NewRouter()
    router.NotFoundHandler = http.HandlerFunc(controllers.Show404Page().Render)
    router.HandleFunc("/", controllers.ShowHomePage().Render)
    router.HandleFunc("/contact", controllers.ShowContactPage().Render)
    router.HandleFunc("/cookietest", controllers.ShowUserCookie).Methods("GET")
    // auth routes.
    auth := router.PathPrefix("/auth").Subrouter()
    auth.HandleFunc("/register", controllers.ParseRegisterForm).Methods("POST")
    auth.HandleFunc("/login", controllers.ParseLoginForm).Methods("POST")
    auth.HandleFunc("/login", controllers.ShowLoginForm().Render).Methods("GET")
    auth.HandleFunc("/register", controllers.ShowRegisterForm().Render).Methods("GET")
    http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), router)

}
