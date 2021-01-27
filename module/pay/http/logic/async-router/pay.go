package async_router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pluto/common"
	"pluto/log"
	"pluto/module/pay/config"
	Proto "pluto/protobuf"
	"pluto/util"
	"strconv"
)

const (
	protoPayCallback			= "PayCallback"

	FmtPayUrl = "consumerId=%s&consumerName=%s&mhtOrderAmt=%s&orderDetail=%s&orderNo=%s&time=%s"
)

func OnPayCallback(c *gin.Context) {

	var (
		err				error
		consumerId		string		// 游戏用户唯一标识
		orderId			string
		orderNo			string		// 平台订单编号
		orderDetail		string		// 订单详情
		mhtOrderAmt		string		// 充值金额(分)
		time			string		// 平台支付时间
		sign			string		// 签名
		url				string
		calSign			string
		uid				uint64
		serverId		uint32
		areaId			uint32
		sid				uint32
		payId			int64
		orderTime		int64
		amountFloat		float64
		amount			int64
	)

	consumerId, err = util.GetParamString(protoPayCallback, "consumerId", c)
	if err != nil {
		return
	}

	orderId, err = util.GetParamString(protoPayCallback, "consumerName", c)
	if err != nil {
		return
	}

	mhtOrderAmt, err = util.GetParamString(protoPayCallback, "mhtOrderAmt", c)
	if err != nil {
		return
	}

	orderDetail, err = util.GetParamString(protoPayCallback, "orderDetail", c)
	if err != nil {
		return
	}

	orderNo, err = util.GetParamString(protoPayCallback, "orderNo", c)
	if err != nil {
		return
	}

	time, err = util.GetParamString(protoPayCallback, "time", c)
	if err != nil {
		return
	}

	sign, err = util.GetParamString(protoPayCallback, "sign", c)
	if err != nil {
		return
	}

	url = fmt.Sprintf(FmtPayUrl, consumerId, orderId, mhtOrderAmt,
		orderDetail, orderNo, time)

	calSign = util.EncryptMd5(url + util.EncryptMd5(config.Config().NASDK.SecretKey))

	if calSign != sign {
		log.Proto(protoPayCallback, false, "invalid sign.")
		return
	}

	uid, err = strconv.ParseUint(consumerId, 0, 64)

	if err != nil {
		log.Proto(protoPayCallback, false, "invalid uid %s.", consumerId)
		return
	}

	serverId = uint32((uid / 100000) % 1000)
	areaId = uint32((uid / 100000) / 1000)
	sid = (areaId << 22) + (2 << 16) + serverId

	log.Debug("ServerID is %d", sid)

	orderTime, err = strconv.ParseInt(time, 0, 64)

	if err != nil {
		log.Proto(protoPayCallback, false, "invalid orderTime %s.", time)
		return
	}

	payId, err = strconv.ParseInt(orderDetail, 0, 64)

	if err != nil {
		log.Proto(protoPayCallback, false, "invalid payId %s.", orderDetail)
		return
	}

	amountFloat, err = strconv.ParseFloat(mhtOrderAmt, 64)

	if err != nil {
		log.Proto(protoPayCallback, false, "invalid amount %s.", mhtOrderAmt)
		return
	}

	amount = int64(amountFloat * 100)

	hd := &common.HttpChData {
		Sid:  sid,
		Data: &Proto.PayCallbackReq{
			ServerID:		sid,
			UserID:        	uid,
			PayID: 		   	int32(payId),
			PurchaseNum:   	1,
			OrderTime:     	orderTime,
			PayAmount:     	int32(amount),
			OrderID:       	orderId,
			SDKOrderID:		orderNo,
		},
	}

	common.HttpCh <- hd

}


