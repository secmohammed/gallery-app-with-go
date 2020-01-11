package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
	"lenslocked.com/config"
	"lenslocked.com/routes"
)

var databaseCredentials = config.GetDatabase()

// User type
type User struct {
	ID    int
	Name  string
	Email string `gorm:"type:varchar(100);unique_index"`
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		databaseCredentials["host"],
		databaseCredentials["port"],
		databaseCredentials["username"],
		databaseCredentials["password"],
		databaseCredentials["database"],
	)
	db, err := gorm.Open("postgres", psqlInfo)
	db.LogMode(true)

	if err != nil {
		panic(err)
	}
	var user User
	db.First(&user, 1)
	log.Fatal(user.Name)
	defer db.Close()

	routes.RegisterRoutes()
}
