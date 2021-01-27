package http

import (
	"github.com/gin-gonic/gin"
	async "pluto/module/list/http/logic/async-router"
	sync "pluto/module/list/http/logic/sync-router"
)

func initRouter() *gin.Engine {

	router := gin.Default()

	// async routers.
	router.POST("/Leave", async.AsyncCb)
	router.POST("/Register", async.AsyncCb)
	router.POST("/ChangeName", async.AsyncCb)

	// sync routers.
	router.POST("/ServerList", sync.OnGetServerList)
	router.POST("/ReloadServerList", sync.OnGMReloadServerList)

	return router
}
