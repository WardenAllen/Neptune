package logic

import (
	"github.com/golang/protobuf/proto"
	"pluto/log"
	Proto "pluto/protobuf"
	"pluto/socket"
)

const (
	protoTCPBannedPlayerAck		= "TCPBannedPlayerAck"
	protoTCPRemovePlayerRankAck	= "TCPRemovePlayerRankAck"
	protoTCPGetOnlinePlayersAck	= "TCPGetOnlinePlayersAck"
	protoTCPNewMailAck			= "TCPNewMailAck"
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

func OnTCPRemovePlayerRankAck(_ *socket.Session, msg *socket.Message) bool {

	ack := &Proto.RemovePlayerRankAck{}
	err := proto.Unmarshal(msg.GetData(), ack)

	if err != nil {
		log.Proto(protoTCPRemovePlayerRankAck, false, "Parse failed.")
		return false
	}

	log.Proto(protoTCPRemovePlayerRankAck, true, "Uid %d.", ack.UserID)

	return true
}

func OnTCPGetOnlinePlayersAck(_ *socket.Session, msg *socket.Message) bool {

	ack := &Proto.GetOnlinePlayersAck{}
	err := proto.Unmarshal(msg.GetData(), ack)

	if err != nil {
		log.Proto(protoTCPGetOnlinePlayersAck, false, "Parse failed.")
		return false
	}

	log.Proto(protoTCPGetOnlinePlayersAck, true, "Players %v.", ack.OnlinePlayers)

	return true
}

func OnTCPNewMailAck(_ *socket.Session, msg *socket.Message) bool {

	ack := &Proto.NewMailAck{}
	err := proto.Unmarshal(msg.GetData(), ack)

	if err != nil {
		log.Proto(protoTCPNewMailAck, false, "Parse failed.")
		return false
	}

	log.Proto(protoTCPNewMailAck, true, ".")

	return true
}
