package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/enzdor/ratings-api/controllers"
    "github.com/enzdor/ratings-api/utils/middlewares"
)

func UsersRouter(r *gin.Engine) {
    users := r.Group("/api/users")

    users.POST("/register", controllers.PostUser)
    users.POST("/login", controllers.LoginUser)
    users.GET("/extend", middlewares.AuthMiddleware(), controllers.ExtendUser)
    users.DELETE("/:user_id", controllers.DeleteUser)
}
