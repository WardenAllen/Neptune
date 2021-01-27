package logic

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"pluto/log"
	"pluto/module/stat/common"
	"pluto/module/stat/config"
	"pluto/module/stat/ta"
	Proto "pluto/protobuf"
	"pluto/socket"
	"pluto/util"
	"strconv"
	"time"
)

const (
	protoTCPOnlineNumReq	= "TCPOnlineNumReq"
	protoTCPPayStatReq		= "TCPPayStatReq"
	protoTCPSyncStatReq		= "TCPSyncStatReq"
)

func OnTCPOnlineNumReq(s *socket.Session, msg *socket.Message) bool {

	var (
		err		error
		req		*Proto.OnlineNumReq
		ack		*Proto.OnlineNumAck
		data	[]byte
		pb		*socket.Message
		channel	uint32
		sid		uint32
		server	string
		oreq	*common.OnlineReq
		reqByte	[]byte
		reqStr	string
		resByte	[]byte
		res		*common.OnlineRes
	)

	req = &Proto.OnlineNumReq{}
	err = proto.Unmarshal(msg.GetData(), req)

	if err != nil {
		log.Proto(protoTCPOnlineNumReq, false, "Parse failed.")
		return false
	}

	// Send ack.

	ack = &Proto.OnlineNumAck{}

	data, err = proto.Marshal(ack)
	if err != nil {
		log.Proto(protoTCPOnlineNumReq, false, "Serialize failed.")
		return false
	}

	pb = socket.NewMessage(uint16(Proto.PLUTO_ID_OnlineNumAckID), data)

	err = s.GetConn().SendMessage(pb)
	if err != nil {
		log.Proto(protoTCPOnlineNumReq, false, "Send failed.")
		return false
	}

	// Post to tapdb.

	channel, _, sid = util.GetServerInfo(req.ServerID)
	server = string(channel) + "-" + string(sid)

	oreq = &common.OnlineReq{
		AppID:   config.Config().TapDB.AppID,
		Onlines: []common.OnlineAry{
			{
				Server:    server,
				Online:    int(req.OnlineNum),
				Timestamp: time.Now().Unix(),
			},
		},
	}

	reqByte, err = json.Marshal(oreq)
	if err != nil {
		log.Proto(protoTCPOnlineNumReq, false, "Build json failed.")
		return false
	}

	reqStr = string(reqByte)

	resByte, _, err = util.HttpPost(config.Config().TapDB.OnlineURL, reqStr)
	if err != nil {
		log.Proto(protoTCPOnlineNumReq, false, "Http post failed.")
		return false
	}

	res = &common.OnlineRes{}

	err = json.Unmarshal(resByte, res)
	if err != nil {
		log.Proto(protoTCPOnlineNumReq, false, "Response parse failed.")
		return false
	}

	if res.Code != 200 {
		log.Proto(protoTCPOnlineNumReq, false, "Response code %d.", res.Code)
		return false
	}

	log.Proto(protoTCPOnlineNumReq, true, "Server %d online num %d.", req.ServerID, req.OnlineNum)

	return true
}

func OnTCPPayStatReq(s *socket.Session, msg *socket.Message) bool {

	var (
		err		error
		req		*Proto.PayStatReq
		ack		*Proto.PayStatAck
		data	[]byte
		pb		*socket.Message
		ip		string
		preq	*common.PayReq
		reqByte	[]byte
		reqStr	string
		status	int
	)

	req = &Proto.PayStatReq{}
	err = proto.Unmarshal(msg.GetData(), req)

	if err != nil {
		log.Proto(protoTCPPayStatReq, false, "Parse failed.")
		return false
	}

	// Send ack.

	ack = &Proto.PayStatAck{}

	data, err = proto.Marshal(ack)
	if err != nil {
		log.Proto(protoTCPPayStatReq, false, "Serialize failed.")
		return false
	}

	pb = socket.NewMessage(uint16(Proto.PLUTO_ID_PayStatAckID), data)

	err = s.GetConn().SendMessage(pb)
	if err != nil {
		log.Proto(protoTCPPayStatReq, false, "Send failed.")
		return false
	}

	// Post to tapdb.

	ip = util.InetNtoA(req.IP)

	preq = &common.PayReq{
		Module:     "GameAnalysis",
		IP:         ip,
		Name:       "charge",
		AppID:      config.Config().TapDB.AppID,
		Account:	req.Account,
		Properties: common.PayProperty{
			OrderID:      req.OrderID,
			Amount:       int(req.Amount),
			ExtraGold:    0,
			CurrencyType: "USD",
			Product:      strconv.FormatInt(int64(req.PayID), 10),
			Payment:      "Google Play",
		},
	}

	reqByte, err = json.Marshal(preq)
	if err != nil {
		log.Proto(protoTCPPayStatReq, false, "Build json failed.")
		return false
	}

	reqStr = string(reqByte)
	reqStr = util.DeleteSpace(reqStr)
	reqStr = util.DeleteCRLF(reqStr)
	reqStr = util.UrlEncode(reqStr)

	_, status, err = util.HttpPost(config.Config().TapDB.PayURL, reqStr)
	if err != nil {
		log.Proto(protoTCPPayStatReq, false, "Http post failed.")
		return false
	}

	if status != 200 {
		log.Proto(protoTCPPayStatReq, false, "Response status code %d.", status)
		return false
	}

	log.Proto(protoTCPPayStatReq, true, "Account %s buy %d pay %d.",
		req.Account, req.PayID, req.Amount)

	return true
}

func OnTCPSyncStatReq(s *socket.Session, msg *socket.Message) bool {

	var (
		err			error
		req			*Proto.TAStatListReq
		ip			string
	)

	req = &Proto.TAStatListReq{}
	err = proto.Unmarshal(msg.GetData(), req)

	if err != nil {
		log.Proto(protoTCPSyncStatReq, false, "Parse failed.")
		return false
	}

	ip = util.InetNtoA(req.Addr)

	for _, evt := range req.EventList {

		err = ta.TrackEvent(req.AccountID, req.DistinctID, ip, req.UserID,
			evt.Id, evt.TimeStamp, evt.IntList, evt.FloatList, evt.StrList)

		if err != nil {
			log.Proto(protoTCPSyncStatReq, false, "ta %d track event failed error %s.", evt.Id, err)
		}

	}

	for _, pty := range req.PropertyList {

		err = ta.TrackProperty(req.AccountID, req.DistinctID, ip, req.UserID,
			pty.Id, pty.IntVal, pty.FloatVal, pty.StrVal, pty.TimeStamp)

		if err != nil {
			log.Proto(protoTCPSyncStatReq, false, "ta %d track property failed error %s.", pty.Id, err)
		}

	}

	return true
}
