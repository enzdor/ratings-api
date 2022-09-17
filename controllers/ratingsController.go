package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/enzdor/ratings-api/utils/models"
    "github.com/enzdor/ratings-api/utils/database"
    "github.com/enzdor/ratings-api/utils/errors"
)


func GetRatings(c *gin.Context) {
    var ratings []models.Rating

    if result := database.Db.Find(&ratings); result.Error != nil {
	err := errors.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, ratings)
}

func GetRatingsByUserID(c *gin.Context) {
    var ratings []models.Rating
    id := c.Param("user_id")

    if result := database.Db.Where("user_id = ?", id).Find(&ratings); result.Error != nil {
	err := errors.NewInternalServerError("user not found in db")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, ratings)
}

func GetRatingByID(c *gin.Context) {
    var rating models.Rating
    id := c.Param("rating_id")

    if result := database.Db.First(&rating, id); result.Error != nil {
	err := errors.NewInternalServerError("rating not found in db")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, rating)
}

func PostRating(c *gin.Context) {
    var rating models.Rating

    if err := c.ShouldBindJSON(&rating); err != nil {
	err := errors.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }

    if result := database.Db.Create(&rating); result.Error != nil {
	err := errors.NewInternalServerError("unable to create rating in db")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, rating)
}

func UpdateRating(c *gin.Context) {
    var rating  models.Rating 
    id := c.Param("rating_id")

    if result := database.Db.First(&rating, id); result.Error != nil {
	err := errors.NewInternalServerError("rating not found in db")
	c.JSON(err.Status, err)
	return
    }

    if err := c.ShouldBindJSON(&rating); err != nil {
	err := errors.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }

    if result := database.Db.Save(&rating); result.Error != nil {
	err := errors.NewInternalServerError("could not update rating in db")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, rating) 
}

func DeleteRating(c *gin.Context) {
    var rating  models.Rating
    id := c.Param("rating_id")

    if result := database.Db.Delete(&rating, id); result.Error != nil {
	err := errors.NewInternalServerError("could not delete rating in db")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, rating) 
}
