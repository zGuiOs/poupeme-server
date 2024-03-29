package router

import (
	"github.com/gin-gonic/gin"
	handler "github.com/zGuiOs/poupeme-server/handlers"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/transaction", handler.ShowOneTransactionHandler)
		v1.GET("/transactions", handler.ShowAllTransactionHandler)
		v1.POST("/transaction", handler.CreateTransactionHandler)
		v1.PUT("/transaction", handler.UpdateTransactionHandler)
		v1.DELETE("/transaction", handler.DeleteTransactionHandler)
	}
}
