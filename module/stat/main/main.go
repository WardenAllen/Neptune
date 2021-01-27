package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"pluto/database"
	"pluto/log"
	"pluto/module/stat/common"
	"pluto/module/stat/config"
	"pluto/module/stat/http"
	_ "pluto/module/stat/ta"
	"pluto/module/stat/tcp"
	"pluto/worker"
)

func main() {

	config.Init()

	// uncomment here when release.
	//log.SetMode(log.ModeRelease)
	gin.SetMode(gin.ReleaseMode)

	// redirect gin's log into file.
	logfile, err := os.OpenFile("./gin-stat.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		log.Error("Can not open log file")
		return
	}
	gin.DefaultWriter = io.MultiWriter(logfile)

	worker.InitDispatcher()

	mc := config.Config().Mysql
	ret := database.InitMysql(mc.Username, mc.Password, mc.Host, mc.Port, mc.DBName)
	if !ret {
		return
	}

	common.Init()

	go tcp.Start()

	http.Init()

	select {}
	
}
