package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "github.com/joho/godotenv"
    "github.com/enzdor/ratings-api/controllers"
    "log"
    "os"
    "fmt"
)

var SecretKey string = os.Getenv("APISECRET")

func StartDatabase() {
    errEnv := godotenv.Load()
    if errEnv != nil {
	log.Fatal(errEnv)
    }
    dbUser := os.Getenv("DBUSER")
    dbPass := os.Getenv("DBPASS")
    dbName := os.Getenv("DBNAME")
    cfg := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbName)

    // Get a database handle.
    var err error

    controllers.Repo.Db, err = gorm.Open(mysql.Open(cfg), &gorm.Config{
	Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        log.Fatal(err)
    }
}
