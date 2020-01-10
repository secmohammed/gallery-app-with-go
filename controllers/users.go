package controllers

import (
    "fmt"
    "net/http"

    "lenslocked.com/resources/views"
)

// ShowRegisterForm function to show the form
func ShowRegisterForm() *Users {
    return &Users{
        NewView: views.NewView("layout", "resources/views/auth/register.gohtml"),
    }
}

//ParseRegisterForm to parse the registration form when submitted.
func ParseRegisterForm(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.PostFormValue("email"))
}

// Users type.
type Users struct {
    NewView *views.View
}

// Render Method to render the parsed view.
func (u *Users) Render(w http.ResponseWriter, r *http.Request) {
    u.NewView.Render(w, nil)
}
