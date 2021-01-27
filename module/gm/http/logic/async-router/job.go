package async_router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pluto/worker"
)

type HttpJob struct {
	Ctx			*gin.Context
	Callback	func (*gin.Context)
}

func (job *HttpJob) Exec (int) error {
	job.Callback(job.Ctx)
	return nil
}

var asyncCbMap = map[string] func (*gin.Context) {
	"/BanPlayer" :		OnBannedPlayer,
	"/RemoveRank" :		OnRemovePlayerRank,
	"/OnlinePlayer" :	OnGetOnlinePlayerInfo,
}

var AsyncCb = func (ctx *gin.Context) {

	f, ok := asyncCbMap[ctx.FullPath()]
	if !ok{
		ctx.JSON(200, gin.H{
			"code" : 100,
			"msg": "invalid path",
		})
	}

	// make a READ-ONLY copy of gin.Context.
	cpy := ctx.Copy()
	// put into job queue.
	worker.JobQueue <- &HttpJob{cpy, f}
	// send response.
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": "success",
	})
}
