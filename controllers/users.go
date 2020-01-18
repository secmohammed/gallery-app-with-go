package controllers

import (
    "fmt"
    "net/http"

    "lenslocked.com/models"
    "lenslocked.com/resources/views"
    "lenslocked.com/utils"
)

// ShowRegisterForm function to show the form
func ShowRegisterForm() *View {
    return &View{
        NewView: views.NewView("auth/register"),
    }
}

// RegisterFormRequest type.
type RegisterFormRequest struct {
    Name     string `schema:"name"`
    Email    string `schema:"email"`
    Password string `schema:"password"`
}

//ParseRegisterForm to parse the registration form when submitted.
func ParseRegisterForm(w http.ResponseWriter, r *http.Request) {
    var form RegisterFormRequest

    utils.Must(utils.ParseForm(r, &form))
    user := models.User{
        Name:     form.Name,
        Email:    form.Email,
        Password: form.Password,
    }
    if err := models.Create(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintln(w, user)
}

// View type.
type View struct {
    NewView *views.View
}

// Render Method to render the parsed view.
func (u *View) Render(w http.ResponseWriter, r *http.Request) {
    u.NewView.Render(w, nil)
}
