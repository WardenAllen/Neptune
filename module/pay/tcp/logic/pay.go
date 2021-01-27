package logic

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"pluto/log"
	Proto "pluto/protobuf"
	"pluto/socket"
)

const (
	protoHTTPPayCallback		= "HTTPPayCallback"
	protoTCPOrder				= "TCPOrder"
	protoTCPPayCallback			= "TCPPayCallback"
)

func OnTCPOrderReq(s *socket.Session, msg *socket.Message) bool {

	var (
		err		error
		req		*Proto.OrderReq
		ack		*Proto.OrderAck
		data	[]byte
		pb		*socket.Message
		orderId string
	)

	req = &Proto.OrderReq{}
	err = proto.Unmarshal(msg.GetData(), req)

	if err != nil {
		log.Proto(protoTCPOrder, false, "Parse failed.")
		return false
	}

	orderId = fmt.Sprintf("%d%d%d", req.OrderTime, req.UserID, req.PayID)

	// TODO: save data here.

	ack = &Proto.OrderAck {
		Error:			0,
		OrderID: 		orderId,
		UserID: 		req.UserID,
		OrderTime: 		req.OrderTime,
		PayID: 			req.PayID,
		PurchaseNum:	req.PurchaseNum,
		PlatformID: 	req.PlatformID,
	}

	data, err = proto.Marshal(ack)

	if err != nil {
		log.Proto(protoTCPOrder, false, "Serialize failed.")
		return false
	}

	pb = socket.NewMessage(uint16(Proto.PLUTO_ID_OrderAckID), data)
	err = s.GetConn().SendMessage(pb)

	if err != nil {
		log.Proto(protoTCPOrder, false, "Send failed.")
		return false
	}

	log.Proto(protoTCPOrder, true, "OrderID %s.", orderId)

	return true
}

func OnHttpPayCallbackReq(s *socket.Session, p *Proto.PayCallbackReq) {

	var (
		err		error
		data	[]byte
		msg		*socket.Message
	)

	data, err = proto.Marshal(p)

	if err != nil {
		log.Proto(protoHTTPPayCallback, false, "Parse failed.")
		return
	}

	msg = socket.NewMessage(uint16(Proto.PLUTO_ID_PayCallbackReqID), data)

	err = s.GetConn().SendMessage(msg)

	if err != nil {
		log.Proto(protoHTTPPayCallback, false, "Send failed.")
		return
	}

	log.Proto(protoHTTPPayCallback, true, "Req uid %d.", p.UserID)

}

func OnTCPPayCallbackAck(_ *socket.Session, msg *socket.Message) bool {

	ack := &Proto.PayCallbackAck{}
	err := proto.Unmarshal(msg.GetData(), ack)

	if err != nil {
		log.Proto(protoTCPPayCallback, false, "Parse failed.")
		return false
	}

	log.Proto(protoTCPPayCallback, true, "")

	return true
}