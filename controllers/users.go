package controllers

import (
    "fmt"
    "net/http"
    "time"

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

//ShowUserCookie function is useed to show the cookies of authenticated user.
func ShowUserCookie(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("remember_token")
    fmt.Println(cookie)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    user, err := models.ByRememberToken(cookie.Value)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintln(w, user)
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
    var form LoginFormRequest
    utils.Must(utils.ParseForm(r, &form))

    user, err := models.Authenticate(form.Email, form.Password)
    if err != nil {
        switch err {
        case models.ErrorNotFound:
            fmt.Fprintln(w, "Invalid Email address")
            break
        case models.ErrorInvalidPassword:
            fmt.Fprintln(w, "Invalid password provided.")
            break
        default:
            fmt.Fprint(w, err.Error())
            break
        }
        return
    }
    err = writeCookieforUser(w, user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/cookietest", http.StatusFound)
}

func writeCookieforUser(w http.ResponseWriter, user *models.User) error {
    if user.RememberToken == "" {
        token, err := utils.RememberToken()
        if err != nil {
            return err
        }
        user.RememberToken = token
        err = models.Update(user)
        if err != nil {
            return err
        }
    }
    expirationTime := time.Now().Add(time.Hour)

    http.SetCookie(w, &http.Cookie{
        Name:     "remember_token",
        Value:    user.RememberToken,
        Expires:  expirationTime,
        HttpOnly: true,
        Secure:   false,
        Path:     "/",
    })
    return nil
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
    err := writeCookieforUser(w, &user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/cookietest", http.StatusFound)
}

// View type.
type View struct {
    NewView *views.View
}

// Render Method to render the parsed view.
func (u *View) Render(w http.ResponseWriter, r *http.Request) {
    u.NewView.Render(w, nil)
}
