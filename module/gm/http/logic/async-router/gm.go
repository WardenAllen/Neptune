package async_router

import (
	"github.com/gin-gonic/gin"
	"pluto/common"
	Proto "pluto/protobuf"
	"pluto/util"
)

const (
	protoBannedPlayer			= "BannedPlayer"
	protoRemovePlayerRank		= "RemovePlayerRank"
)

func OnBannedPlayer(c *gin.Context) {

	var (
		err			error
		sid			uint32
		uid			uint64
		banTime		int
	)

	sid, err = util.GetParamUInt32(protoBannedPlayer, "sid", c)
	if err != nil {
		return
	}

	uid, err = util.GetParamUInt64(protoBannedPlayer, "uid", c)
	if err != nil {
		return
	}

	banTime, err = util.GetParamInt(protoBannedPlayer, "banTime", c)
	if err != nil {
		return
	}

	hd := &common.HttpChData {
		Sid:  sid,
		Data: &Proto.BannedPlayerReq {
			UserID:     uid,
			Bannedtime: int32(banTime),
		},
	}

	common.HttpCh <- hd

}

func OnRemovePlayerRank(c *gin.Context) {

	var (
		err			error
		sid			uint32
		uid			uint64
	)

	sid, err = util.GetParamUInt32(protoRemovePlayerRank, "sid", c)
	if err != nil {
		return
	}

	uid, err = util.GetParamUInt64(protoRemovePlayerRank, "uid", c)
	if err != nil {
		return
	}

	hd := &common.HttpChData {
		Sid:  sid,
		Data: &Proto.RemovePlayerRankReq {
			UserID:     uid,
		},
	}

	common.HttpCh <- hd

}

func OnGetOnlinePlayerInfo(*gin.Context) {

	hd := &common.HttpChData {
		Sid:  0,
		Data: &Proto.GetOnlinePlayersReq{},
	}

	common.HttpCh <- hd

}

