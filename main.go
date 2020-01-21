package main

import (

    // load dotenv variables
    "fmt"

    _ "github.com/joho/godotenv/autoload"

    // load migrations
    _ "lenslocked.com/migrations"
    "lenslocked.com/models"
    "lenslocked.com/routes"
)

func main() {
    user := models.User{
        Name:          "asomeone",
        Email:         "onetwothree@gmail.com",
        Password:      "jondasd",
        RememberToken: "abc1234",
    }
    err := models.Create(&user)
    if err != nil {
        panic(err)
    }
    user2, err := models.ByRememberToken("abc1234")
    if err != nil {
        panic(err)
    }
    fmt.Println(user2)
    routes.RegisterRoutes()
}
