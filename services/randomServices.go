package services

import (
    "gorm.io/gorm"
    "os"
)

var SecretKey string = os.Getenv("APISECRET")

var Db *gorm.DB
