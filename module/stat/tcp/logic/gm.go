package logic

import (
	"github.com/golang/protobuf/proto"
	"pluto/log"
	Proto "pluto/protobuf"
	"pluto/socket"
)

const (
	protoTCPBannedPlayerAck		= "TCPBannedPlayerAck"
)

func OnTCPBannedPlayerAck(_ *socket.Session, msg *socket.Message) bool {

	ack := &Proto.BannedPlayerAck{}
	err := proto.Unmarshal(msg.GetData(), ack)

	if err != nil {
		log.Proto(protoTCPBannedPlayerAck, false, "Parse failed.")
		return false
	}

	log.Proto(protoTCPBannedPlayerAck, true, "Uid %d time %d.", ack.UserID, ack.Bannedtime)

	return true
}
