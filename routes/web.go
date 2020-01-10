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

    router.HandleFunc("/auth/register", controllers.ShowRegisterForm().Render)
    http.ListenAndServe(":3000", router)

}
