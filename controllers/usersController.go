package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "time"
    "strconv"
    "github.com/golang-jwt/jwt"
    "github.com/enzdor/ratings-api/models"
    "github.com/enzdor/ratings-api/services"
    "github.com/enzdor/ratings-api/utils"
    "golang.org/x/crypto/bcrypt"
)


func GetUserByID(c *gin.Context) {
    var user models.User
    id := c.Param("user_id")

    if result := services.Db.First(&user, id); result.Error != nil {
	err := utils.NewBadRequestError("not found in db")
	c.JSON(err.Status, err)
	return
    }

    c.JSON(http.StatusOK, user)
}

func PostUser(c *gin.Context) {
    var user models.User

    if err:= c.ShouldBindJSON(&user); err != nil {
	err := utils.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }

    newPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
    if err != nil {
	err := utils.NewInternalServerError("could not encrypt password")
	c.JSON(err.Status, err)
	return
    }
    user.Password = string(newPassword)

    if result := services.Db.Create(&user); result.Error != nil {
	err := utils.NewInternalServerError("unable to create user in db")
	c.JSON(err.Status, err)
	return
    }

    claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	ExpiresAt: time.Now().Add(time.Hour).Unix(),
	Issuer: strconv.Itoa(user.User_id),
    })

    token, err := claims.SignedString([]byte(services.SecretKey))
    if err != nil {
	err := utils.NewInternalServerError("unable to create token")
	c.JSON(err.Status, err)
	return
    }

    c.SetCookie("jwt", token, 100, "/", "localhost", false, true)
    c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
    var user  models.User
    id := c.Param("user_id")

    if result := services.Db.Delete(&user, id); result.Error != nil {
	err := utils.NewInternalServerError("unable to delete user")
	c.JSON(err.Status, err)
	return
    }

    c.JSON(http.StatusOK, user) 
}

func LoginUser(c *gin.Context) {
    var reqUser models.User
    var dbUser models.User

    if err:= c.ShouldBindJSON(&reqUser); err != nil {
	err := utils.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }
    if result := services.Db.Where("email = ?", reqUser.Email).First(&dbUser); result.Error != nil {
	err := utils.NewInternalServerError("user not found in db")
	c.JSON(err.Status, err)
	return
    }
    if dbUser.User_id == 0 {
	err := utils.NewInternalServerError("user not found in db")
	c.JSON(err.Status, err)
	return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(reqUser.Password)); err != nil {
	err := utils.NewInternalServerError("incorrect password")
	c.JSON(err.Status, err)
	return
    }

    claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	ExpiresAt: time.Now().Add(time.Hour).Unix(),
	Issuer: strconv.Itoa(dbUser.User_id),
    })

    token, err := claims.SignedString([]byte(services.SecretKey))
    if err != nil {
	err := utils.NewInternalServerError("unable to create token")
	c.JSON(err.Status, err)
	return
    }

    c.SetCookie("jwt", token, 100, "/", "localhost", false, true)
    c.JSON(http.StatusOK, token)
}

func LogoutUser(c *gin.Context) {
    c.SetCookie("jwt", "", -1, "", "", false, true)
    c.JSON(http.StatusOK, "logout succesful")
}
