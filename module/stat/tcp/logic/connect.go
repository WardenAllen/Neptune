package logic

import (
	"github.com/golang/protobuf/proto"
	"pluto/log"
	Proto "pluto/protobuf"
	"pluto/socket"
)

const (
	protoTCPConncetReq	= "TCPConncetReq"
)

func OnTCPConncetReq(s *socket.Session, msg *socket.Message) bool {

	var (
		err		error
		req		*Proto.PlutoConnectReq
		ack		*Proto.PlutoConnectAck
		data	[]byte
		pb		*socket.Message
	)

	req = &Proto.PlutoConnectReq{}
	err = proto.Unmarshal(msg.GetData(), req)

	if err != nil {
		log.Proto(protoTCPConncetReq, false, "Parse failed.")
		return false
	}

	s.GetService().RegServer(req.ServerID, s)

	ack = &Proto.PlutoConnectAck{
		Code:     1,
		ServerID: 1 << 16 + 8 << 8 + 1,
	}

	data, err = proto.Marshal(ack)

	if err != nil {
		log.Proto(protoTCPConncetReq, false, "Serialize failed.")
		return false
	}

	pb = socket.NewMessage(uint16(Proto.PLUTO_ID_ConnectAckID), data)
	err = s.GetConn().SendMessage(pb)

	if err != nil {
		log.Proto(protoTCPConncetReq, false, "Send failed.")
		return false
	}

	log.Proto(protoTCPConncetReq, true, "Connect sid %d.", req.ServerID)

	return true
}