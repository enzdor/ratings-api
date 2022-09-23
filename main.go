package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "github.com/joho/godotenv"
    "github.com/enzdor/ratings-api/routers"
    "github.com/enzdor/ratings-api/utils/database"
    "log"
    "os"
    "fmt"
    "time"
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
    database.Db, err = gorm.Open(mysql.Open(cfg), &gorm.Config{
	Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        log.Fatal(err)
    }

    router := gin.Default()
    routers.RatingsRouter(router)
    routers.UsersRouter(router)

    router.Use(cors.New(cors.Config{
	AllowOrigins:	[]string{"http://localhost:3000"},
	AllowMethods:	[]string{"POST", "GET", "PATCH", "PUT"},
	AllowHeaders:	[]string{"Content-Type"},
	ExposeHeaders:	[]string{"Content-Length"},
	AllowCredentials: true,
	AllowOriginFunc: func(origin string) bool {
	    return origin == "http://localhost:3000"
	},
	MaxAge:		12 * time.Hour,
    }))
    router.Run("localhost:8080")
}
