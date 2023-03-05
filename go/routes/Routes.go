package routes

import (
	"web-backend/go/controller"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("card")
	{
		userGroup.POST("/cards", controller.CreateCard)
		userGroup.GET("/cards", controller.GetCardList)
	}
	return r
}
