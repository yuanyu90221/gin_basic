package main

import (
	"github.com/yuanyu90221/gin_basic/router"
)

func main() {
	ginRouter := router.SetupRouter()
	_ = ginRouter.Run()
}
