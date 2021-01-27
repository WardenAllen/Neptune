package http

import (
	"github.com/gin-gonic/gin"
	async "pluto/module/stat/http/logic/async-router"
)

func initRouter() *gin.Engine {

	router := gin.Default()

	// async routers.
	router.POST("/", async.AsyncCb)

	return router
}
