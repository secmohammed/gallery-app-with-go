package models

import (
    "errors"

    "github.com/jinzhu/gorm"
    "golang.org/x/crypto/bcrypt"
    "lenslocked.com/utils"

    // load the database connection.
    _ "lenslocked.com/utils"
)

// db variable to get the database connection.
var db *gorm.DB = utils.GetDatabaseConnection()

// ErrorNotFound is returned when a resource cannot be found.
var ErrorNotFound = errors.New("model: resource not found")

//ErrorInvaildID will be thrown in case of the id is invalid or equal to zero.
var ErrorInvaildID = errors.New("models: ID provided was invalid")

// User type
type User struct {
    gorm.Model
    Name     string
    Email    string `gorm:"not null;type:varchar(100);unique_index"`
    Password string `gorm:"not null;type:varchar(100);"`
}

//Create function is used to create a users record
func Create(user *User) error {
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedBytes)
    return db.Create(user).Error
}

// Update will update the provided user with all of the data passed through the user object.
func Update(user *User) error {
    return db.Save(user).Error
}

// ByID will look up the users using the given id.
func ByID(id uint) (*User, error) {
    var user User
    query := db.Where("id = ?", id)
    err := first(query, &user)
    return &user, err
}

//Delete function will delete the user with the provided id.
func Delete(id uint) error {
    // gorm will delete all of the records if the id equals to zero.
    if id == 0 {
        return ErrorInvaildID
    }
    // Gorm delete needs the primary key with the reference of the object to understand which table we are deleting from.
    user := User{
        Model: gorm.Model{
            ID: id,
        },
    }
    return db.Delete(&user).Error
}

// ByEail function will lok up the users using the given email.
func ByEmail(email string) (*User, error) {
    var user User
    query := db.Where(&user, "email = ?", email)
    err := first(query, &user)
    return &user, err
}

// first will quer using the provided database query and it will get the first item returned and place it into dst, if nothing is returned/found in the query, it will return error not found
func first(db *gorm.DB, dst interface{}) error {
    err := db.First(dst).Error
    switch err {
    case nil:
        return nil
    case gorm.ErrRecordNotFound:
        return ErrorNotFound
    default:
        return err
    }

}
