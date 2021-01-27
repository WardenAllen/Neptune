package sync_router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pluto/common"
	"pluto/log"
	Proto "pluto/protobuf"
)

const (
	protoSendMail				= "SendMail"
)

func OnSendMail(c *gin.Context) {

	type awardItem struct {
		ItemID		int32		`form:"itemID" json:"itemID"`
		ItemNum		int32		`form:"itemNum" json:"itemNum"`
	}

	type newMailReq struct {
		Sid			uint32		`form:"sid" json:"sid"`
		Title       string		`form:"title" json:"title"`
		Text        string		`form:"text" json:"text"`
		DeleteType  int32		`form:"deleteType" json:"deleteType"`
		DeleteParam int32		`form:"deleteParam" json:"deleteParam"`
		AccountType int32		`form:"accountType" json:"accountType"`
		UserIDs     []uint64	`form:"userIDs" json:"userIDs"`
		Awards      []awardItem	`form:"awards" json:"awards"`
	}

	var (
		err			error
		req			newMailReq
		uids		[]uint64
		awards		[]*Proto.AwardItem
	)

	defer func() {
		if err := recover(); err != nil {
			log.Error("Send mail panic %v.", err)
		}
	}()

	req = newMailReq{}
	err = c.ShouldBindJSON(&req)
	if err != nil {

		log.Proto(protoSendMail, false, "Parse body error %s.", err.Error())

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})

		return
	}

	for i := 0; i < len(req.UserIDs); i++ {
		uids = append(uids, req.UserIDs[i])
	}

	for i := 0; i < len(req.Awards); i++ {

		award := Proto.AwardItem{
			ItemID:  req.Awards[i].ItemID,
			ItemNum: req.Awards[i].ItemNum,
		}

		awards = append(awards, &award)
	}

	hd := &common.HttpChData {
		Sid:  req.Sid,
		Data:&Proto.NewMailReq{
			Title:       req.Title,
			Text:        req.Text,
			DeleteType:  req.DeleteType,
			DeleteParam: req.DeleteParam,
			AccountType: req.AccountType,
			UserIDs:     uids,
			Awards:      awards,
		},
	}

	common.HttpCh <- hd

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": "success",
	})

}
