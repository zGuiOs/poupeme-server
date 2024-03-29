package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/transaction", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Api initialized",
			})
		})
		v1.POST("/transaction", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Api initialized",
			})
		})
		v1.DELETE("/transaction", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Api initialized",
			})
		})
		v1.PUT("/transaction", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Api initialized",
			})
		})
		v1.GET("/transactions", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Api initialized",
			})
		})
	}
}
