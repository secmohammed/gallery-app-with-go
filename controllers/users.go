package controllers

import (
    "fmt"
    "net/http"

    "lenslocked.com/models"
    "lenslocked.com/resources/views"
    "lenslocked.com/utils"
)

// ShowLoginForm function is used to show the login form.
func ShowLoginForm() *View {
    return &View{
        NewView: views.NewView("auth/login"),
    }
}

// ShowRegisterForm function to show the form
func ShowRegisterForm() *View {
    return &View{
        NewView: views.NewView("auth/register"),
    }
}

//LoginFormRequest type.
type LoginFormRequest struct {
    Email    string `schema:"email"`
    Password string `schema:"password"`
}

// RegisterFormRequest type.
type RegisterFormRequest struct {
    Name     string `schema:"name"`
    Email    string `schema:"email"`
    Password string `schema:"password"`
}

//ParseLoginForm to parse the login form when submitted.
func ParseLoginForm(w http.ResponseWriter, r *http.Request) {
    form := LoginFormRequest{}
    utils.Must(utils.ParseForm(r, &form))
    // user := models.User{
    //     Email:    form.Email,
    //     Password: form.Password,
    // }
    fmt.Fprintln(w, form)
    // if user, err := models.Login(&user); err != nil {
    //     return
    // }
    // log.Fatal(user)
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
