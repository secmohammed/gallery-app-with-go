package models

import "github.com/jinzhu/gorm"

// User type
type User struct {
    gorm.Model
    ID    int
    Name  string
    Email string `gorm:"type:varchar(100);unique_index"`
}
