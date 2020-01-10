package views

import "html/template"

// NewView function to create a new view by parsing passed templates.
func NewView(layout string, files ...string) *View {
    files = append(files,
        "views/layouts/footer.gohtml",
        "views/layouts/master.gohtml",
    )
    t, err := template.ParseFiles(files...)
    if err != nil {
        panic(err)
    }
    return &View{
        Template: t,
        Layout:   layout,
    }
}

// View struct, to create a new object file and parse the view files.
type View struct {
    Template *template.Template
    Layout   string
}
