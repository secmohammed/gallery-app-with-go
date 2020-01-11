package migrations

import (
    "lenslocked.com/models"
    "lenslocked.com/utils"
)

func init() {
    utils.GetDatabaseConnection().AutoMigrate(models.User{})
}
