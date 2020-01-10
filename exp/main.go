package main

import (
    "html/template"
    "os"
)

// User type
type User struct {
    Name  string
    Slice []string
}

func main() {
    t, err := template.ParseFiles("hello.gohtml")
    if err != nil {
        panic(err.Error())
    }
    user := User{
        Name:  "John Smith",
        Slice: []string{"a", "b", "c"},
    }
    err = t.Execute(os.Stdout, user)
    if err != nil {
        panic(err)
    }
}
