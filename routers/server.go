package routers

import (
    "os"
    "time"
    "github.com/gin-contrib/cors" 
    "github.com/enzdor/ratings-api/utils/middlewares"
    "github.com/enzdor/ratings-api/controllers"
    "github.com/gin-gonic/gin"
)

var SecretKey string = os.Getenv("APISECRETKEY")

func StartRouters() {
    router := gin.Default()
    ratingsRouter(router)
    usersRouter(router)

    router.Use(cors.New(cors.Config{
	AllowOrigins:     []string{"http://localhost:3000"},
	AllowMethods:     []string{"POST", "GET", "PATCH", "PUT"},
	AllowHeaders:     []string{"Content-Type"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
	AllowOriginFunc: func(origin string) bool {
	    return origin == "http://localhost:3000"
	},
	MaxAge: 12 * time.Hour,
    }))
    router.Run("localhost:8080")
}

func usersRouter(r *gin.Engine) {
    users := r.Group("/api/users")

    users.POST("/register", controllers.Repo.PostUser)
    users.POST("/login", controllers.Repo.LoginUser)
    users.GET("/extend", middlewares.AuthMiddleware(), controllers.Repo.ExtendUser)
    users.DELETE("/:user_id", controllers.Repo.DeleteUser)
}

func ratingsRouter(r *gin.Engine) {
    ratings := r.Group("/api/ratings")

    ratings.Use(middlewares.AuthMiddleware())

    ratings.GET("/:rating_id", controllers.Repo.GetRatingByID)
    ratings.GET("/user", controllers.Repo.GetRatingsByUserID)
    ratings.POST("/", controllers.Repo.PostRating)
    ratings.POST("/search", controllers.Repo.SearchRatingsByUserID)
    ratings.PATCH("/:rating_id", controllers.Repo.UpdateRating)
    ratings.DELETE("/:rating_id", controllers.Repo.DeleteRating)
}
