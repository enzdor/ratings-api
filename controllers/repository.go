package controllers

import (
    "os"
    "gorm.io/gorm"
)

var SecretKey string = os.Getenv("APISECRET")

type Repository struct {
    Db *gorm.DB
}

var Repo Repository
