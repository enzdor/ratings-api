package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/enzdor/ratings-api/models"
    "github.com/enzdor/ratings-api/services"
    "github.com/enzdor/ratings-api/utils"
)


func GetRatings(c *gin.Context) {
    var ratings []models.Rating

    if result := services.Db.Find(&ratings); result.Error != nil {
	err := utils.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, ratings)
}

func GetRatingsByUserID(c *gin.Context) {
    var ratings []models.Rating
    id := c.Param("user_id")

    if result := services.Db.Where("user_id = ?", id).Find(&ratings); result.Error != nil {
	err := utils.NewInternalServerError("user not found in db")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, ratings)
}

func GetRatingByID(c *gin.Context) {
    var rating models.Rating
    id := c.Param("rating_id")

    if result := services.Db.First(&rating, id); result.Error != nil {
	err := utils.NewInternalServerError("rating not found in db")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, rating)
}

func PostRating(c *gin.Context) {
    var rating models.Rating

    if err := c.ShouldBindJSON(&rating); err != nil {
	err := utils.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }

    if result := services.Db.Create(&rating); result.Error != nil {
	err := utils.NewInternalServerError("unable to create rating in db")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, rating)
}

func UpdateRating(c *gin.Context) {
    var rating  models.Rating 
    id := c.Param("rating_id")

    if result := services.Db.First(&rating, id); result.Error != nil {
	err := utils.NewInternalServerError("rating not found in db")
	c.JSON(err.Status, err)
	return
    }

    if err := c.ShouldBindJSON(&rating); err != nil {
	err := utils.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }

    if result := services.Db.Save(&rating); result.Error != nil {
	err := utils.NewInternalServerError("could not update rating in db")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, rating) 
}

func DeleteRating(c *gin.Context) {
    var rating  models.Rating
    id := c.Param("rating_id")

    if result := services.Db.Delete(&rating, id); result.Error != nil {
	err := utils.NewInternalServerError("could not delete rating in db")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, rating) 
}
