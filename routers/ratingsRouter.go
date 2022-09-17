package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/enzdor/ratings-api/controllers"
    "github.com/enzdor/ratings-api/utils/middlewares"
)

func RatingsRouter(r *gin.Engine) {
    ratings := r.Group("/api/ratings")

    ratings.GET("/", controllers.GetRatings)
    ratings.GET("/:rating_id", loggedornot.AuthMiddleware(), controllers.GetRatingByID)
    ratings.GET("/user/:user_id", controllers.GetRatingsByUserID)
    ratings.POST("/", controllers.PostRating)
    ratings.PATCH("/", controllers.UpdateRating)
    ratings.DELETE("/", controllers.DeleteRating)
}
