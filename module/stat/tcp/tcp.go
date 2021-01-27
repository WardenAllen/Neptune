package tcp

import (
	"pluto/log"
	"pluto/module/stat/config"
	Proto "pluto/protobuf"
	"pluto/socket"
)

func init() {
}

func Start() {

	ss, err := socket.NewSocketService(config.Config().Server.TcpHost)
	if err != nil {
		return
	}

	// ss.SetHeartBeat(5*time.Second, 30*time.Second)

	ss.RegMessageHandler(HandleMessage)
	ss.RegConnectHandler(HandleConnect)
	ss.RegDisconnectHandler(HandleDisconnect)
	ss.RegHttpHandler(HandleHttpMsg)

	//go NewClientConnect()
	//
	//timer := time.NewTimer(time.Second * 5)
	//go func() {
	//	<-timer.C
	//
	//	req := &Proto.RemovePlayerRankReq{
	//		UserID: 111222,
	//		Uuid:   "",
	//	}
	//
	//	common.HttpCh <- req
	//
	//}()

	log.Info("TCP Server listening at " + config.Config().Server.TcpHost + ".")

	ss.Start()

}

func HandleMessage(s *socket.Session, msg *socket.Message) {

	handler, ok := protoMap[Proto.PLUTO_ID(msg.GetID())]

	if !ok {
		return
	}

	handler(s, msg)

}

func HandleDisconnect(s *socket.Session, err error) {

	log.Warn(s.GetConn().GetName() + " lost connect.")

}

func HandleConnect(s *socket.Session) {

	log.Info(s.GetConn().GetName() + " connected.")

	//req := &Proto.ConnectAck {
	//	ServerID: 1 << 16 + 7 << 8 + 1,
	//}
	//protoData, _ := proto.Marshal(req)
	//msg := socket.NewMessage(Proto.MSGID_Connect_Ack_ID, protoData)
	//_ = s.GetConn().SendMessage(msg)

}

