package views

import (
    "html/template"
    "path/filepath"
)

var (
    layoutDir         = "views/layouts/"
    templateExtension = ".gohtml"
)

// NewView function to create a new view by parsing passed templates.
// when function is first letter uppercase it's already exported, if we don't want to export it, we name it normally.
func NewView(layout string, files ...string) *View {

    files = append(files,
        layoutFiles(layoutDir)...,
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

// layout files returns a slice of strings representing
// the layout files used in our application.
func layoutFiles(layoutDir string) []string {
    files, err := filepath.Glob(layoutDir + "*" + templateExtension)
    if err != nil {
        panic(err)
    }
    return files
}
