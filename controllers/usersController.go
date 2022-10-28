package controllers

import (
    "fmt"
    "net/http"
    "strconv"
    "github.com/golang-jwt/jwt"
    "github.com/gin-gonic/gin"
    "github.com/enzdor/ratings-api/utils/models"
    "github.com/enzdor/ratings-api/utils/helpers"
    "github.com/enzdor/ratings-api/utils/errors"
    "golang.org/x/crypto/bcrypt"
)


func (r *Repository) PostUser(c *gin.Context) {
    var user, checkUser models.User

    if err:= c.ShouldBindJSON(&user); err != nil {
	err := errors.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }

    if result := r.Db.Where("email = ?", user.Email).First(&checkUser); result.Error == nil {
	err := errors.NewBadRequestError("user already exists")
	c.JSON(err.Status, err)
	return
    }

    newPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
    if err != nil {
	err := errors.NewInternalServerError("could not encrypt password")
	c.JSON(err.Status, err)
	return
    }
    user.Password = string(newPassword)

    if result := r.Db.Create(&user); result.Error != nil {
	err := errors.NewInternalServerError("unable to create user in db")
	c.JSON(err.Status, err)
	return
    }

    token, tokenerr := helpers.CreateToken(user.User_id) 
    if tokenerr != nil {
	c.JSON(tokenerr.Status, tokenerr)
	return
    }

    c.JSON(http.StatusOK, token)
}

func (r *Repository) DeleteUser(c *gin.Context) {
    var user  models.User
    id := c.Param("user_id")

    if result := r.Db.Delete(&user, id); result.Error != nil {
	err := errors.NewInternalServerError("unable to delete user")
	c.JSON(err.Status, err)
	return
    }

    c.JSON(http.StatusOK, user) 
}

func (r *Repository) LoginUser(c *gin.Context) {
    var reqUser models.User
    var dbUser models.User

    if err:= c.ShouldBindJSON(&reqUser); err != nil {
	err := errors.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }
    if result := r.Db.Where("email = ?", reqUser.Email).First(&dbUser); result.Error != nil {
	err := errors.NewInternalServerError("user not found in db")
	c.JSON(err.Status, err)
	return
    }
    if dbUser.User_id == 0 {
	err := errors.NewInternalServerError("user not found in db")
	c.JSON(err.Status, err)
	return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(reqUser.Password)); err != nil {
	err := errors.NewInternalServerError("incorrect password")
	c.JSON(err.Status, err)
	return
    }

    token, err := helpers.CreateToken(dbUser.User_id); if err != nil {
	c.JSON(err.Status, err)
	return
    }

    c.JSON(http.StatusOK, token)
}

func (r *Repository) ExtendUser(c *gin.Context) {
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
	return []byte(SecretKey), nil
    })
    if err != nil {
	err := errors.NewInternalServerError("error parsing header")
	c.AbortWithStatusJSON(err.Status, err)
	return
    }

    claims := token.Claims.(*jwt.StandardClaims)

    issuer, err := strconv.Atoi(claims.Issuer)
    if err != nil {
	err := errors.NewInternalServerError("unable to parse stirng")
	c.JSON(err.Status, err)
	return
    }

    newToken, tokenerr := helpers.CreateToken(issuer) 
    if tokenerr != nil {
	c.JSON(tokenerr.Status, tokenerr)
	return
    }

    c.JSON(http.StatusOK, newToken)
}
