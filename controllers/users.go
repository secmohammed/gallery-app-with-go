package controllers

import (
    "fmt"
    "net/http"

    "github.com/gorilla/schema"
    "lenslocked.com/resources/views"
)

// ShowRegisterForm function to show the form
func ShowRegisterForm() *View {
    return &View{
        NewView: views.NewView("layout", "resources/views/auth/register.gohtml"),
    }
}

// RegisterFormRequest type.
type RegisterFormRequest struct {
    Email    string `schema:"email"`
    Password string `schema:"password"`
}

//ParseRegisterForm to parse the registration form when submitted.
func ParseRegisterForm(w http.ResponseWriter, r *http.Request) {
    // parseForm must be called in order to fill the postForm with the data coming from the input form data.
    if err := r.ParseForm(); err != nil {
        panic(err)
    }
    decoder := schema.NewDecoder()
    var form RegisterFormRequest

    if err := decoder.Decode(&form, r.PostForm); err != nil {
        panic(err)
    }
    fmt.Fprintln(w, form)
}

// View type.
type View struct {
    NewView *views.View
}

// Render Method to render the parsed view.
func (u *View) Render(w http.ResponseWriter, r *http.Request) {
    u.NewView.Render(w, nil)
}
