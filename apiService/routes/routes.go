package routes

import (
	"example.com/controllers"
	"github.com/gin-gonic/gin"
)

func Init() (route *gin.Engine) {
	route = gin.Default()
	route.GET("/blocks", controllers.GetBlocks)
	route.GET("/blocks/:id", controllers.GetBlock)
	route.GET("/transaction/:txHash", controllers.GetTransaction)
	return route
}
