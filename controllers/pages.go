package controllers

import (
    "net/http"

    "lenslocked.com/resources/views"
)

// ShowHomePage function to show the page.
func ShowHomePage() *Pages {
    return &Pages{
        NewView: views.NewView("layout", "home"),
    }
}

//ShowContactPage function is used to show the contact page.
func ShowContactPage() *Pages {
    return &Pages{
        NewView: views.NewView("layout", "contact"),
    }
}

// Show404Page function is used to be shown whenever the route isn't found.
func Show404Page() *Pages {
    return &Pages{
        NewView: views.NewView("layout", "404"),
    }
}

// Pages type.
type Pages struct {
    NewView *views.View
}

// Render Method to render the parsed view.
func (page *Pages) Render(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    page.NewView.Render(w, nil)
}
