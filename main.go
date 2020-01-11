package main

import (
	"database/sql"
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	"lenslocked.com/config"
	"lenslocked.com/routes"
)

var databaseCredentials = config.GetDatabase()

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		databaseCredentials["host"],
		databaseCredentials["port"],
		databaseCredentials["username"],
		databaseCredentials["password"],
		databaseCredentials["database"],
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	routes.RegisterRoutes()
}
