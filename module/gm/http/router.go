package http

import (
	"github.com/gin-gonic/gin"
	async "pluto/module/gm/http/logic/async-router"
	sync "pluto/module/gm/http/logic/sync-router"
)

func initRouter() *gin.Engine {

	router := gin.Default()

	// async routers.
	router.POST("/BanPlayer", async.AsyncCb)
	router.POST("/RemoveRank", async.AsyncCb)
	router.POST("/OnlinePlayer", async.AsyncCb)
	router.POST("/Mail", sync.OnSendMail)

	return router
}
