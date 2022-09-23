package helpers

import (
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/enzdor/ratings-api/utils/errors"
    "github.com/enzdor/ratings-api/utils/database"
    "github.com/golang-jwt/jwt"
)

func GetUserId(c *gin.Context) (issuer int){
    header := c.GetHeader("jwt-token")

    if header == "" {
	err := errors.NewBadRequestError("invalid token string haha")
	c.AbortWithStatusJSON(err.Status, err)
    }

    token, err := jwt.ParseWithClaims(header, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
	return []byte(database.SecretKey), nil
    })
    if err != nil {
	err := errors.NewInternalServerError("error parsing header")
	c.AbortWithStatusJSON(err.Status, err)
    }

    claims := token.Claims.(*jwt.StandardClaims)
    issuer64, err := strconv.ParseInt(claims.Issuer, 10, 64)
    if err != nil {
	err := errors.NewInternalServerError("user id should be number")
	c.AbortWithStatusJSON(err.Status, err)
    }

    issuer = int(issuer64)

    return issuer
}
