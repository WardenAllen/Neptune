package logic

import (
	"github.com/golang/protobuf/proto"
	"pluto/log"
	Proto "pluto/protobuf"
	"pluto/socket"
)

const (
	protoHTTPBannedPlayer		= "HTTPBannedPlayer"
	protoHTTPRemovePlayerRank	= "HttpRemovePlayerRank"
	protoHTTPGetOnlinePlayers	= "HttpGetOnlinePlayers"
	protoHTTPNewMail			= "HTTPNewMail"
)

func OnHttpBannedPlayer(s *socket.Session, p *Proto.BannedPlayerReq) {

	var (
		err		error
		data	[]byte
		msg		*socket.Message
	)

	data, err = proto.Marshal(p)

	if err != nil {
		log.Proto(protoHTTPBannedPlayer, false, "Parse failed.")
		return
	}

	msg = socket.NewMessage(uint16(Proto.PLUTO_ID_BannedPlayerReqID), data)

	err = s.GetConn().SendMessage(msg)

	if err != nil {
		log.Proto(protoHTTPBannedPlayer, false, "Send failed.")
		return
	}

	log.Proto(protoHTTPBannedPlayer, true, "Req uid %d time %d.", p.UserID, p.Bannedtime)

}

func OnHttpRemovePlayerRank(s *socket.Session, p *Proto.RemovePlayerRankReq) {

	var (
		err		error
		data	[]byte
		msg		*socket.Message
	)

	data, err = proto.Marshal(p)

	if err != nil {
		log.Proto(protoHTTPRemovePlayerRank, false, "Parse failed.")
		return
	}

	msg = socket.NewMessage(uint16(Proto.PLUTO_ID_RemovePlayerRankReqID), data)

	err = s.GetConn().SendMessage(msg)

	if err != nil {
		log.Proto(protoHTTPRemovePlayerRank, false, "Send failed.")
		return
	}

	log.Proto(protoHTTPRemovePlayerRank, true, "Req uid %d.", p.UserID)

}

func OnHttpGetOnlinePlayers(s *socket.Session, p *Proto.GetOnlinePlayersReq) {

	var (
		err		error
		data	[]byte
		msg		*socket.Message
	)

	data, err = proto.Marshal(p)

	if err != nil {
		log.Proto(protoHTTPGetOnlinePlayers, false, "Parse failed.")
		return
	}

	msg = socket.NewMessage(uint16(Proto.PLUTO_ID_GetOnlinePlayersReqID), data)

	err = s.GetConn().SendMessage(msg)

	if err != nil {
		log.Proto(protoHTTPGetOnlinePlayers, false, "Send failed.")
		return
	}

	log.Proto(protoHTTPGetOnlinePlayers, true, "Req.")

}

func OnHttpNewMailReq(s *socket.Session, p *Proto.NewMailReq) {

	var (
		err		error
		data	[]byte
		msg		*socket.Message
	)

	data, err = proto.Marshal(p)

	if err != nil {
		log.Proto(protoHTTPNewMail, false, "Parse failed.")
		return
	}

	msg = socket.NewMessage(uint16(Proto.PLUTO_ID_NewMailReqID), data)

	err = s.GetConn().SendMessage(msg)

	if err != nil {
		log.Proto(protoHTTPNewMail, false, "Send failed.")
		return
	}

	log.Proto(protoHTTPNewMail, true,
		"Req title %s text %s award %s uids %v.", p.Title, p.Text, p.Awards, p.UserIDs)

}