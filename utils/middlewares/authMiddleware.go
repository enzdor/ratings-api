package middlewares

import (
    "fmt"
    "time"
    "strconv"
    "github.com/enzdor/ratings-api/utils/models"
    "github.com/enzdor/ratings-api/utils/errors"
    "github.com/enzdor/ratings-api/utils/database"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
    return func (c *gin.Context) {
	header := c.GetHeader("jwt-token")

	if header == "" {
	    err := errors.NewBadRequestError("invalid token string here")
	    c.AbortWithStatusJSON(err.Status, err)
	    return
	}

	token, err := jwt.ParseWithClaims(header, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
	    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		err := errors.NewBadRequestError("unexpected singing method")
		c.AbortWithStatusJSON(err.Status, err)
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	    }
	    return []byte(database.SecretKey), nil
	})
	if err != nil {
	    err := errors.NewInternalServerError("error parsing header")
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

	if claims.ExpiresAt < time.Now().Unix() {
	    err := errors.NewBadRequestError("token expired")
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
