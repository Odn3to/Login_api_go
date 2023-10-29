package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"column:username"`
    Password string `gorm:"column:password"`
    Token    string `gorm:"column:token"`
}
