package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "github.com/joho/godotenv"
    "github.com/enzdor/ratings-api/routers"
    "github.com/enzdor/ratings-api/services"
    "log"
    "os"
    "fmt"
)

func main() {
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
    services.Db, err = gorm.Open(mysql.Open(cfg), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"https://localhost:8080", "https://ratings-gray.vercel.app", "https://hoppscotch.io"}

    router := gin.Default()
    routers.RatingsRouter(router)
    routers.UsersRouter(router)

    router.Run("localhost:8080")
}


