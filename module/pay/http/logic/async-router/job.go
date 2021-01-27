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
	"/PayCallback" : OnPayCallback,
}

var AsyncCb = func (ctx *gin.Context) {

	f, ok := asyncCbMap[ctx.FullPath()]
	if !ok {
		ctx.String(http.StatusOK, "FAIL")
	}

	// make a READ-ONLY copy of gin.Context.
	cpy := ctx.Copy()
	// put into job queue.
	worker.JobQueue <- &HttpJob{cpy, f}

	// TODO: Insert order info to mysql.

	// send response.
	ctx.String(http.StatusOK, "SUCCESS")
}
