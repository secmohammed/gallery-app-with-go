package controllers

import (
    "net/http"

    "lenslocked.com/resources/views"
)

// ShowRegisterForm function to show the form
func ShowRegisterForm() *Users {
    return &Users{
        NewView: views.NewView("layout", "resources/views/auth/register.gohtml"),
    }
}

// Users type.
type Users struct {
    NewView *views.View
}

// Render Method to render the parsed view.
func (u *Users) Render(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    u.NewView.Render(w, nil)
}
