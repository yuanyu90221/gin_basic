package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	md "github.com/yuanyu90221/gin_basic/middle"
)

func SetupRouter() *gin.Engine {
	initRouter := gin.Default()
	initRouter.Use(md.Logger())
	// get router
	initRouter.GET("/", func(context *gin.Context) {
		// context.String(http.StatusOK, "hello gin")
		middle := context.MustGet("middle").(string)
		context.JSON(http.StatusOK, gin.H{
			"message": "hello gin",
			"middle":  middle,
		})
	})
	return initRouter
}
