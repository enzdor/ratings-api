package loggedornot

import (
    "strconv"
    "github.com/enzdor/ratings-api/utils/models"
    "github.com/enzdor/ratings-api/utils/errors"
    "github.com/enzdor/ratings-api/utils/database"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
    return func (c *gin.Context) {
	cookie := c.GetHeader("jwt-token")

	if cookie == "" {
	    err := errors.NewBadRequestError("invalid token string")
	    c.AbortWithStatusJSON(err.Status, err)
	    return
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
	    return []byte(database.SecretKey), nil
	})
	if err != nil {
	    err := errors.NewInternalServerError("error parsing cookie")
	    c.AbortWithStatusJSON(err.Status, err)
	    return
	}

	claims := token.Claims.(*jwt.StandardClaims)
	issuer, err := strconv.ParseInt(claims.Issuer, 10, 64)
	if err != nil {
	    err := errors.NewInternalServerError("user id should be number")
	    c.AbortWithStatusJSON(err.Status, err)
	    return
	}

	var user models.User

	if result := database.Db.First(&user, issuer); result.Error != nil {
	    err := errors.NewBadRequestError("could not find user in db")
	    c.AbortWithStatusJSON(err.Status, err)
	    return
	}

	c.Next()
    }
}
