package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/enzdor/ratings-api/controllers"
    "github.com/enzdor/ratings-api/utils/middlewares"
)

func RatingsRouter(r *gin.Engine) {
    ratings := r.Group("/api/ratings")

    ratings.Use(middlewares.AuthMiddleware())

    ratings.GET("/:rating_id", controllers.GetRatingByID)
    ratings.GET("/user", controllers.GetRatingsByUserID)
    ratings.POST("/", controllers.PostRating)
    ratings.POST("/search", controllers.SearchRatingsByUserID)
    ratings.PATCH("/:rating_id", controllers.UpdateRating)
    ratings.DELETE("/:rating_id", controllers.DeleteRating)
}
