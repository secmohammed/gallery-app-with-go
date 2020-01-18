package main

import (
    "fmt"
    // load dotenv variables
    _ "github.com/joho/godotenv/autoload"
    // load migrations
    _ "lenslocked.com/migrations"

    "lenslocked.com/routes"
)

func main() {
    routes.RegisterRoutes()
    fmt.Println("starting the server on :3000")
}
