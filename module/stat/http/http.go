package http

import (
	"pluto/log"
	"pluto/module/stat/config"
)

func Init()  {

	router := initRouter()

	log.Info("Router init success.")

	log.Info("HTTP Server listening at " + config.Config().Server.HttpHost + ".")

	err := router.Run(config.Config().Server.HttpHost)

	if err != nil {
		log.Error("Http server init failed %s!", err.Error())
		return
	}

}