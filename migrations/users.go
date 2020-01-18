package migrations

import (
    "lenslocked.com/models"
    "lenslocked.com/utils"
)

func init() {
    utils.GetDatabaseConnection().AutoMigrate(models.User{})
}

//refresh function is used to take the tables down form the database and refresh it
func Refresh() {
    utils.GetDatabaseConnection().DropTableIfExists(&models.User{})
    utils.GetDatabaseConnection().AutoMigrate(&models.User{})
}
