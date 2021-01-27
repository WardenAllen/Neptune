package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pluto/log"
)

type HttpServer struct {
	server		*http.Server
}

func (hs *HttpServer) Start (addr string, router *gin.Engine) {

	hs.server = &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		// service connections
		if err := hs.server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {

			log.Error("HttpServer listen: %s", err)

		}
	}()

	//// Wait for interrupt signal to gracefully shutdown the server with
	//// a timeout of 5 seconds.
	//quit := make(chan os.Signal)
	//signal.Notify(quit, syscall.SIGUSR1)
	//<-quit
	//
	//log.Info("Shutdown http server ...")
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//
	//defer cancel()
	//if err := hs.server.Shutdown(ctx); err != nil {
	//	log.Error("Shutdown http server failed %s.", err)
	//}
	//
	//log.Info("Shutdown http server success.")

}