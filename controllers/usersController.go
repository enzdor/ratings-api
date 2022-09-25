package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "time"
    "strconv"
    "github.com/golang-jwt/jwt"
    "github.com/enzdor/ratings-api/utils/models"
    "github.com/enzdor/ratings-api/utils/database"
    "github.com/enzdor/ratings-api/utils/errors"
    "golang.org/x/crypto/bcrypt"
)


func PostUser(c *gin.Context) {
    var user, checkUser models.User

    if err:= c.ShouldBindJSON(&user); err != nil {
	err := errors.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }

    if result := database.Db.Where("email = ?", user.Email).First(&checkUser); result.Error == nil {
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

    if result := database.Db.Create(&user); result.Error != nil {
	err := errors.NewInternalServerError("unable to create user in db")
	c.JSON(err.Status, err)
	return
    }

    claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	ExpiresAt: time.Now().Add(time.Hour).Unix(),
	Issuer: strconv.Itoa(user.User_id),
    })

    token, err := claims.SignedString([]byte(database.SecretKey))
    if err != nil {
	err := errors.NewInternalServerError("unable to create token")
	c.JSON(err.Status, err)
	return
    }

    c.JSON(http.StatusOK, token)
}

func DeleteUser(c *gin.Context) {
    var user  models.User
    id := c.Param("user_id")

    if result := database.Db.Delete(&user, id); result.Error != nil {
	err := errors.NewInternalServerError("unable to delete user")
	c.JSON(err.Status, err)
	return
    }

    c.JSON(http.StatusOK, user) 
}

func LoginUser(c *gin.Context) {
    var reqUser models.User
    var dbUser models.User

    if err:= c.ShouldBindJSON(&reqUser); err != nil {
	err := errors.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }
    if result := database.Db.Where("email = ?", reqUser.Email).First(&dbUser); result.Error != nil {
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

    claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
	Issuer: strconv.Itoa(dbUser.User_id),
    })

    token, err := claims.SignedString([]byte(database.SecretKey))
    if err != nil {
	err := errors.NewInternalServerError("unable to create token")
	c.JSON(err.Status, err)
	return
    }

    c.JSON(http.StatusOK, token)
}

func LogoutUser(c *gin.Context) {
    c.SetCookie("jwt-token", "", -1, "", "", true, true)
    c.JSON(http.StatusOK, "logout succesful")
}
