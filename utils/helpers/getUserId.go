package helpers

import (
    "fmt"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/enzdor/ratings-api/utils/errors"
    "github.com/enzdor/ratings-api/utils/models"
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


type TypeQuery struct {
	Rating_id	int `json:"rating_id" gorm:"primaryKey;unique;notNull"`
	Name		string `json:"name" gorm:"notNull;size:255"`
	Entry_type	string `json:"entry_type" gorm:"notNull;size:255"`
	Rating		string `json:"rating" gorm:"notNull"`
	Consumed	string `json:"consumed" gorm:"notNull"`
	User_id		int `json:"user_id" gorm:"notNull;foreignKey:user_id;references:user_id;constraint:OnUpdate,OnDelete"`
}

func (searchQuery *TypeQuery) CreateSearchQuery(searchRating *models.SearchRating) {
    searchQuery.Name = fmt.Sprintf("%s%s%s", "%", searchRating.Name, "%")
    searchQuery.Entry_type = fmt.Sprintf("%s%s%s", "%", searchRating.Entry_type, "%")
    searchQuery.User_id = searchRating.User_id

    if searchRating.Rating == -1 {
	searchQuery.Rating = "%%"
    } else {
	searchQuery.Rating = strconv.Itoa(searchRating.Rating)
    }

    if searchRating.Consumed == -1 {
	searchQuery.Consumed = "%%"
    } else {
	fmt.Println(searchQuery.Consumed)
	searchQuery.Consumed = strconv.Itoa(searchRating.Consumed)
    }
}
