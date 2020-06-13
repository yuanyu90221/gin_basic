package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	md "github.com/yuanyu90221/gin_basic/middle"
)

func SetupRouter() *gin.Engine {
	initRouter := gin.Default()
	initRouter.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
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
