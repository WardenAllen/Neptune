package tcp

import (
	"pluto/module/pay/tcp/logic"
	Proto "pluto/protobuf"
	"pluto/socket"
)

var protoMap = map[Proto.PLUTO_ID] func (*socket.Session, *socket.Message) bool {
	Proto.PLUTO_ID_ConnectReqID:		logic.OnTCPConncetReq,
	Proto.PLUTO_ID_OrderReqID:			logic.OnTCPOrderReq,
	Proto.PLUTO_ID_PayCallbackAckID:	logic.OnTCPPayCallbackAck,
}

func HandleHttpMsg(s *socket.Session, p interface{}) {

	switch pb := p.(type) {

	case *Proto.PayCallbackReq:
		logic.OnHttpPayCallbackReq(s, pb)

	default:
		return

	}
}