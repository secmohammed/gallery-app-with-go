package main

import (
    "fmt"
    "os"

    // load dotenv variables
    _ "github.com/joho/godotenv/autoload"
    // load migrations
    _ "lenslocked.com/migrations"
    "lenslocked.com/routes"
)

func main() {
    fmt.Printf("server started on :%s", os.Getenv("APP_PORT"))
    routes.RegisterRoutes()
}
