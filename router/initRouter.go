package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	initRouter := gin.Default()
	// get router
	initRouter.GET("/", func(context *gin.Context) {
		// context.String(http.StatusOK, "hello gin")
		context.JSON(http.StatusOK, gin.H{
			"message": "hello gin",
		})
	})
	return initRouter
}
