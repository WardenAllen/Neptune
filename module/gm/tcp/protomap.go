package tcp

import (
	"pluto/module/gm/tcp/logic"
	Proto "pluto/protobuf"
	"pluto/socket"
)

var protoMap = map[Proto.PLUTO_ID] func (*socket.Session, *socket.Message) bool {
	Proto.PLUTO_ID_ConnectReqID:			logic.OnTCPConncetReq,
	Proto.PLUTO_ID_BannedPlayerAckID:		logic.OnTCPBannedPlayerAck,
	Proto.PLUTO_ID_RemovePlayerRankAckID:	logic.OnTCPRemovePlayerRankAck,
	Proto.PLUTO_ID_GetOnlinePlayersAckID:	logic.OnTCPGetOnlinePlayersAck,
	Proto.PLUTO_ID_NewMailAckID:			logic.OnTCPNewMailAck,
	Proto.PLUTO_ID_CDKeyReqID:				logic.OnTCPCdkeyReq,
}

func HandleHttpMsg(s *socket.Session, p interface{}) {

	switch pb := p.(type) {

	case *Proto.RemovePlayerRankReq:
		logic.OnHttpRemovePlayerRank(s, pb)

	case *Proto.BannedPlayerReq:
		logic.OnHttpBannedPlayer(s, pb)

	case *Proto.GetOnlinePlayersReq:
		logic.OnHttpGetOnlinePlayers(s, pb)

	case *Proto.NewMailReq:
		logic.OnHttpNewMailReq(s, pb)

	default:
		return

	}
}
