package views

import "html/template"

// NewView function to create a new view by parsing passed templates.
func NewView(files ...string) (*View, error) {
    files = append(files, "views/layouts/footer.gohtml")
    t, err := template.ParseFiles(files...)
    if err != nil {
        return nil, err
    }
    return &View{
        Template: t,
    }, err
}

// View struct, to create a new object file and parse the view files.
type View struct {
    Template *template.Template
}
