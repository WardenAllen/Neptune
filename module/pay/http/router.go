package http

import (
	async "pluto/module/pay/http/logic/async-router"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {

	router := gin.Default()

	// async routers.
	router.GET("/PayCallback", async.AsyncCb)

	return router
}
