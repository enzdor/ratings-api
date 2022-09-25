package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/enzdor/ratings-api/utils/models"
    "github.com/enzdor/ratings-api/utils/database"
    "github.com/enzdor/ratings-api/utils/errors"
    "github.com/enzdor/ratings-api/utils/helpers"
)


func GetRatingsByUserID(c *gin.Context) {
    var ratings []models.Rating
    id := helpers.GetUserId(c)

    if result := database.Db.Where("user_id = ?", id).Find(&ratings); result.Error != nil {
	err := errors.NewInternalServerError("user not found in db")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, ratings)
}

func SearchRatingsByUserID(c *gin.Context) {
    var searchRating models.SearchRating
    var ratings []models.Rating
    var searchQuery helpers.TypeQuery 

    if err := c.ShouldBindJSON(&searchRating); err != nil {
	err := errors.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }

    searchQuery.CreateSearchQuery(&searchRating)
    id := helpers.GetUserId(c)
    searchQuery.User_id = id

    if result := database.Db.Where(
	    "name LIKE ? AND entry_type LIKE ? AND rating LIKE ? AND consumed LIKE ? AND user_id = ?",
	    searchQuery.Name,
	    searchQuery.Entry_type,
	    searchQuery.Rating,
	    searchQuery.Consumed,
	    searchQuery.User_id,
	).Find(&ratings); result.Error != nil {
	err := errors.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, ratings)
}

func GetRatingByID(c *gin.Context) {
    var rating models.Rating
    issuer := helpers.GetUserId(c)
    id := c.Param("rating_id")

    if result := database.Db.First(&rating, id); result.Error != nil {
	err := errors.NewInternalServerError("rating not found in db")
	c.JSON(err.Status, err)
	return
    }

    if rating.User_id != issuer {
	err := errors.NewBadRequestError("rating does not belong to user")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, rating)
}

func PostRating(c *gin.Context) {
    issuer := helpers.GetUserId(c)

    var rating models.Rating

    if err := c.ShouldBindJSON(&rating); err != nil {
	err := errors.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }

    rating.User_id = int(issuer)

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
    issuer := helpers.GetUserId(c)

    if result := database.Db.First(&rating, id); result.Error != nil {
	err := errors.NewInternalServerError("rating not found in db")
	c.JSON(err.Status, err)
	return
    }

    if rating.User_id != issuer {
	err := errors.NewBadRequestError("rating does not belong to user")
	c.JSON(err.Status, err)
	return
    }

    if err := c.ShouldBindJSON(&rating); err != nil {
	err := errors.NewInternalServerError("invalid json body")
	c.JSON(err.Status, err)
	return
    }

    rating.User_id = int(issuer)

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
    issuer := helpers.GetUserId(c)

    if result := database.Db.First(&rating, id); result.Error != nil {
	err := errors.NewInternalServerError("rating not found in db")
	c.JSON(err.Status, err)
	return
    }

    if rating.User_id != issuer {
	err := errors.NewBadRequestError("rating does not belong to user")
	c.JSON(err.Status, err)
	return
    }

    if result := database.Db.Delete(&rating, id); result.Error != nil {
	err := errors.NewInternalServerError("could not delete rating in db")
	c.JSON(err.Status, err)
	return
    }

    c.IndentedJSON(http.StatusOK, rating) 
}
