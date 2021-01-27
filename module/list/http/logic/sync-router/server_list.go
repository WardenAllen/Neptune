package sync_router

import (
	"github.com/gin-gonic/gin"
	"pluto/log"
	"pluto/module/list/common"
	"pluto/module/list/config"
	"pluto/util"
)

type ListRespData struct {
	Conf		[]config.ServerInfo
	Notice		string
	List		[]*common.User
}

type ListResp struct {
	Code		int
	Msg			string
	Data		ListRespData
}

const (
	protoGetServerList	= "GetServerList"
)

func OnGetServerList(c *gin.Context) {

	var (
		data	ListRespData
		m		map[int]*common.User
		ok		bool
		err		error
		account	string
		channel	int
		code	= 0
		msg		= "success"
	)

	account, err = util.GetParamStringCanBeNil(protoGetServerList, "account", c)

	if err != nil {
		code, msg = 1, "invalid parameters"
		goto RESP
	}

	channel, err = util.GetParamInt(protoGetServerList, "channel", c)

	if err != nil {
		code, msg = 1, "invalid parameters"
		goto RESP
	}

	common.UserRLock()
	defer common.UserRUnlock()

	data.Conf = config.Config().ServerList.List
	data.Notice = config.GetNotice(channel)

	if len(account) != 0 {

		m, ok = common.UsersMap[account]

		if ok {

			for _, v := range m {

				if v.Channel != channel {
					continue
				}

				data.List = append(data.List, v)

			}

		}
	}

	if data.List == nil {
		data.List = make([]*common.User, 0)
	}

	log.Proto(protoGetServerList, true, "User %s channel %d.", account, channel)

RESP:
	if code == 0 {
		c.JSON(200, gin.H{
			"code":	code,
			"msg":	msg,
			"data":	data,
		})
	} else {
		c.JSON(200, gin.H{
			"code":	code,
			"msg":	msg,
		})
	}

}
