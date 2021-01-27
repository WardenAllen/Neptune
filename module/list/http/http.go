package http

import (
	"pluto/log"
	"pluto/module/list/config"
)

func Init()  {

	router := initRouter()

	log.Info("HTTP Server listening at " + config.Config().Server.HttpHost + ".")

	err := router.Run(config.Config().Server.HttpHost)

	if err != nil {
		log.Error("Http server init failed err %s!", err)
	}
}