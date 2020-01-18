package models
import(
    "testing"
    "fmt"
    // setting up gorm with postgres connection
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "github.com/jinzhu/gorm"

)


//env, err = godotenv.Unmarshal("APP_ENV=testing")
//err = godotenv.Write(env, "../.env")
func TestCreateUser(t *testing.T) {
    db, err := testingUserService()
    if err != nil {
        t.Fatal(err)
    }
    user := User{
        Name: "Mohammd Osama",
        Password: "helloworld",
        Email: "mohammeadosama@ieee.org",

    }
    err = Create(&user)
    if err != nil {
        t.Fatal(err)
    }
    if user.ID == 0 {
        t.Errorf("Expected ID > 0, Received %d", user.ID)
    }
}