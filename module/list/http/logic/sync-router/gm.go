package sync_router

import (
	"github.com/gin-gonic/gin"
	"pluto/log"
	"pluto/module/list/config"
	"pluto/util"
)

const (
	protoGMReloadServerList	= "GMReloadServerList"
)

func OnGMReloadServerList(c *gin.Context) {

	var (
		code	= 0
		msg		= "success"
		err		error
		key		string
	)

	key, err = util.GetParamString(protoGMReloadServerList, "key", c)

	if err != nil || key != "WardenAllenPluto" {

		log.Proto(protoGMReloadServerList, false,
			"Invalid key %s from %s %s.", key, c.Request.RemoteAddr, c.Request.URL)

		code, msg = 2, "illegal request"
		goto RESP

	}

	// reload configs.
	config.LoadConfig()

	log.Proto(protoGMReloadServerList, true, ".")

RESP:
	c.JSON(200, gin.H{
		"code":	code,
		"msg":	msg,
	})

}