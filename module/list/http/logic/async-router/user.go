package async_router

import (
	"github.com/gin-gonic/gin"
	"pluto/log"
	"pluto/module/list/common"
	"pluto/util"
	"time"
)

const (
	protoUserRegister		= "UserRegister"
	protoUserLeaveGame		= "UserLeaveGame"
	protoUserChangeRoleName	= "UserChangeRoleName"
)

func OnUserRegister(idx int, c *gin.Context) {

	var (
		err			error
		uid			uint64
		account		string
		rolename	string
		sid			int
		channel		int
		icon		int
	)

	uid, err = util.GetParamUInt64(protoUserRegister, "uid", c)

	if err != nil {
		return
	}

	account, err = util.GetParamString(protoUserRegister, "account", c)

	if err != nil {
		return
	}

	rolename, err = util.GetParamString(protoUserRegister, "rolename", c)

	if err != nil {
		return
	}

	sid, err = util.GetParamInt(protoUserRegister, "sid", c)

	if err != nil {
		return
	}

	channel, err = util.GetParamInt(protoUserRegister, "channel", c)

	if err != nil {
		return
	}

	icon, err = util.GetParamInt(protoUserRegister, "icon", c)

	if err != nil {
		return
	}

	// TODO: whether channel is valid

	common.UserWLock()
	defer common.UserWUnlock()

	t := time.Now()

	u := &common.User {
		UserID:		uid,
		Account:	account,
		RoleName: 	rolename,
		ServerID:	sid,
		Channel: 	channel,
		Level: 		1,
		Icon: 		icon,
		RegTime:	t.Unix(),
		ActiveTime: t.Unix(),
	}

	if !common.InsertUser(u) {
		log.Proto(protoUserRegister, false, "User %s channel %d server %d insert failed.", account, channel, sid)
		return
	}

	err = u.Insert(idx)

	if err != nil {
		log.Proto(protoUserRegister, false, "User %s channel %d server %d insert to db failed.", account, channel, sid)
	} else {
		log.Proto(protoUserRegister, true, "User %s channel %d server %d.", account, channel, sid)
	}

}

func OnUserLeaveGame(idx int, c *gin.Context) {

	var (
		err			error
		account		string
		sid			int
		level		int
		icon		int
	)

	account, err = util.GetParamString(protoUserLeaveGame, "account", c)

	if err != nil {
		return
	}

	sid, err = util.GetParamInt(protoUserLeaveGame, "sid", c)

	if err != nil {
		return
	}

	level, err = util.GetParamInt(protoUserLeaveGame, "level", c)

	if err != nil {
		return
	}

	icon, err = util.GetParamInt(protoUserLeaveGame, "icon", c)

	if err != nil {
		return
	}

	common.UserWLock()
	defer common.UserWUnlock()

	if !common.IsUserExist(account, sid, false) {
		log.Proto(protoUserLeaveGame, false, "User %s at server%d not exist.", account, sid)
		return
	}

	common.UsersMap[account][sid].UpdateLevel(idx, level, icon)

	log.Proto(protoUserLeaveGame, true, "User %s at server%d.", account, sid)

}

func OnUserChangeRoleName(idx int, c *gin.Context) {

	var (
		err			error
		account		string
		sid			int
		rolename	string
	)

	account, err = util.GetParamString(protoUserChangeRoleName, "account", c)

	if err != nil {
		return
	}

	sid, err = util.GetParamInt(protoUserChangeRoleName, "sid", c)

	if err != nil {
		return
	}

	rolename, err = util.GetParamString(protoUserChangeRoleName, "rolename", c)

	if err != nil {
		return
	}

	common.UserWLock()
	defer common.UserWUnlock()

	if !common.IsUserExist(account, sid, false) {
		log.Proto(protoUserChangeRoleName, false, "User %s at server%d not exist.", account, sid)
		return
	}

	common.UsersMap[account][sid].UpdateRoleName(idx, rolename)

	log.Proto(protoUserChangeRoleName, true, "User %s at server%d.", account, sid)

}