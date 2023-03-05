package controller

import (
	"net/http"
	"web-backend/go/entity"
	"web-backend/go/service"

	"github.com/gin-gonic/gin"
)

func CreateCard(c *gin.Context) {

	var card entity.Card
	c.BindJSON(&card)
	err := service.CreateCard(&card)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": card,
		})
	}
}

func GetCardList(c *gin.Context) {
	cardList, err := service.GetAllCards()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": cardList,
		})
	}
}
