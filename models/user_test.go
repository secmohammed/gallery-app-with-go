package models

import (
    "testing"
    // setting up gorm with postgres connection
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

//env, err = godotenv.Unmarshal("APP_ENV=testing")
//err = godotenv.Write(env, "../.env")
func TestCreateUser(t *testing.T) {
    user := User{
        Name:     "Mohammd Osama",
        Password: "helloworld",
        Email:    "mohammeadosama@ieee.org",
    }
    if user.ID == 0 {
        t.Errorf("Expected ID > 0, Received %d", user.ID)
    }
}
