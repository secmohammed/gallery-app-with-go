package config

import (
    "os"
)
var db = map[string]string{
    "port":       os.Getenv("DB_PORT"),
    "database":   os.Getenv("DB_DATABASE"),
    "host":       os.Getenv("DB_HOST"),
    "password":   os.Getenv("DB_PASSWORD"),
    "username":   os.Getenv("DB_USERNAME"),
    "connection": os.Getenv("DB_CONNECTION"),
}
var dbTesting = map[string]string{
    "port":       os.Getenv("DB_PORT"),
    "database":   "fresh_test",
    "host":       os.Getenv("DB_HOST"),
    "password":   os.Getenv("DB_PASSWORD"),
    "username":   os.Getenv("DB_USERNAME"),
    "connection": os.Getenv("DB_CONNECTION"),
}
// GetDatabase function is used to retrieve the database credentials from the env
func GetDatabase() map[string]string {
    return db
}
// GetTestDatabase function is used to retrieve the database credentials from the env
func GetTestDatabase() map[string]string {
    return dbTesting
}