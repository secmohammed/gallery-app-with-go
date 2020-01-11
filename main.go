package main

import (
	"log"
	// load dotenv variables
	_ "github.com/joho/godotenv/autoload"
	// load migrations
	_ "lenslocked.com/migrations"
	"lenslocked.com/models"
	"lenslocked.com/routes"
	"lenslocked.com/utils"
	_ "lenslocked.com/utils"
)

func main() {
	db := utils.GetDatabaseConnection()
	var user models.User
	db.First(&user, 1)
	log.Fatal(user.Name)
	routes.RegisterRoutes()
}
