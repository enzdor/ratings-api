package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/enzdor/ratings-api/controllers"
)

func RatingsRouter(r *gin.Engine) {
    ratings := r.Group("/ratings")


    ratings.GET("/", controllers.GetRatings)
    ratings.GET("/:rating_id", controllers.GetRatingByID)
    ratings.GET("/user/:user_id", controllers.GetRatingsByUserID)
    ratings.POST("/", controllers.PostRating)
    ratings.PATCH("/", controllers.UpdateRating)
    ratings.DELETE("/", controllers.DeleteRating)

}
