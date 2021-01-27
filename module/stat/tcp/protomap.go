package tcp

import (
	"pluto/module/stat/tcp/logic"
	Proto "pluto/protobuf"
	"pluto/socket"
)

var protoMap = map[Proto.PLUTO_ID] func (*socket.Session, *socket.Message) bool {
	Proto.PLUTO_ID_ConnectReqID:			logic.OnTCPConncetReq,
	Proto.PLUTO_ID_OnlineNumReqID:			logic.OnTCPOnlineNumReq,
	Proto.PLUTO_ID_PayStatReqID:			logic.OnTCPPayStatReq,
	Proto.PLUTO_ID_SyncStatReqID: 			logic.OnTCPSyncStatReq,
}

func HandleHttpMsg(s *socket.Session, p interface{}) {

}
