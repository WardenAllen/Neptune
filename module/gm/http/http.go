package http

import (
	"pluto/http"
	"pluto/log"
	"pluto/module/gm/config"
)

func Init()  {

	router := initRouter()

	log.Info("Router init success.")

	hs := &http.HttpServer{}

	log.Info("HTTP Server listening at " + config.Config().Server.HttpHost + ".")

	hs.Start(config.Config().Server.HttpHost, router)

}