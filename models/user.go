package models

import (
    "errors"
    "os"

    "github.com/jinzhu/gorm"
    "golang.org/x/crypto/bcrypt"
    "lenslocked.com/utils"

    // load the database connection.
    _ "lenslocked.com/utils"
)

var (
    // db variable to get the database connection.
    db *gorm.DB = utils.GetDatabaseConnection()
    //HmacSecretKey is a variable that defines the secret key needed for hmac
    HmacSecretKey = os.Getenv("HMAC_SECRET_KEY")
    //hmac variable to get the hmac hash we created
    hmac = utils.NewHMAC(HmacSecretKey)

    // ErrorNotFound is returned when a resource cannot be found.
    ErrorNotFound = errors.New("model: resource not found")

    //ErrorInvaildID will be thrown in case of the id is invalid or equal to zero.
    ErrorInvaildID = errors.New("models: ID provided was invalid")

    //ErrorInvalidPassword will be thrown in case of password mismatch
    ErrorInvalidPassword = errors.New("models: incorrect password provided")
)

// User type
type User struct {
    gorm.Model
    Name          string
    Email         string `gorm:"not null;type:varchar(100);unique_index"`
    Password      string `gorm:"not null;type:varchar(100);"`
    RememberToken string `gorm:"nullable;unique_index"`
}

//Create function is used to create a users record
func Create(user *User) error {
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedBytes)
    user, err = generateRememberTokenFor(user)
    if err != nil {
        return err
    }
    return db.Create(user).Error
}
func generateRememberTokenFor(user *User) (*User, error) {
    if user.RememberToken == "" {
        token, err := utils.RememberToken()
        if err != nil {
            return nil, err
        }
        user.RememberToken = token
    }
    user.RememberToken = hmac.Hash(user.RememberToken)
    return user, nil
}

// Update will update the provided user with all of the data passed through the user object.
func Update(user *User) error {
    user, err := generateRememberTokenFor(user)
    if err != nil {
        return err
    }
    return db.Save(user).Error
}

//ByRememberToken function is used to fetch the user by the provided token.
func ByRememberToken(token string) (*User, error) {
    var user User
    err := first(db.Where("remember_token = ?", token), &user)
    if err != nil {
        return nil, err
    }
    return &user, nil
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

// Authenticate function is used to authorize user throughout his credentials.
func Authenticate(email, password string) (*User, error) {
    user, err := ByEmail(email)
    if err != nil {
        return nil, err
    }
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        switch err {
        case bcrypt.ErrMismatchedHashAndPassword:
            return nil, ErrorInvalidPassword
        default:
            return nil, err
        }
    }
    return user, nil
}

// ByEmail function will lok up the users using the given email.
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
