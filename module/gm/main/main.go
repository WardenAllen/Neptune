package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"pluto/database"
	"pluto/log"
	"pluto/module/gm/config"
	"pluto/module/gm/http"
	"pluto/module/gm/tcp"
	"pluto/worker"
)

func main() {

	/* 1. init configs. */
	config.Init()

	/* 2. start tcp server. */
	worker.InitDispatcher()

	mc := config.Config().Mysql
	ret := database.InitMysql(mc.Username, mc.Password, mc.Host, mc.Port, mc.DBName)
	if !ret {
		return
	}

	go tcp.Start()

	/* 3. start http server. */
	gin.SetMode(gin.ReleaseMode)
	// log.SetMode(log.ModeRelease)

	// redirect gin's log into file.
	logfile, err := os.OpenFile("./gin-gm.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		log.Error("Can not open log file")
		return
	}
	gin.DefaultWriter = io.MultiWriter(logfile)

	http.Init()

	select {}
	
}
