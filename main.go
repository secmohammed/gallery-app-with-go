package main

import (

    // load dotenv variables

    _ "github.com/joho/godotenv/autoload"
    // load migrations
    _ "lenslocked.com/migrations"

    "lenslocked.com/models"
    "lenslocked.com/routes"
)

func main() {
    user := models.User{
        Name: "Mohammd Osama",
        Password: "helloworld",
        Email: "mohammeadosama@ieee.org",

    }
    if err:= models.Create(&user); err != nil {
    }
    user.Email = "someone@gmail.com"
    if err:= models.Update(&user); err != nil {
    }
    if err:= models.Delete(user.ID); err != nil {
        panic(err)
    }
    routes.RegisterRoutes()

}
