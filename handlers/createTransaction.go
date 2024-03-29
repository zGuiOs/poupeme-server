package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransactionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Api initialized",
	})
}
