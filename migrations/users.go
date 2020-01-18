package migrations

import (
    "lenslocked.com/models"
    "lenslocked.com/utils"
)

var database = utils.GetDatabaseConnection()

func init() {
    database.AutoMigrate(models.User{})
}

//Refresh function is used to take the tables down form the database and refresh it
func Refresh() {
    database.DropTableIfExists(&models.User{})
    database.AutoMigrate(&models.User{})
}
