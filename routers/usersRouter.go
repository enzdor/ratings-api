package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/enzdor/ratings-api/controllers"
)

func UsersRouter(r *gin.Engine) {
    users := r.Group("/api/users")

    users.POST("/register", controllers.PostUser)
    users.POST("/login", controllers.LoginUser)
    users.GET("/:user_id", controllers.GetUserByID)
    users.GET("/logout", controllers.LogoutUser)
    users.DELETE("/:user_id", controllers.DeleteUser)
}
