package common

import "net"

type Socket struct {
	Conn 	net.Conn
}

type SendBuf struct {
	Buf			[]byte
	ServerType	int32
}

type HttpChData struct {
	Sid			uint32
	Data		interface{}
}

var HttpCh chan *HttpChData

func Init() {
	HttpCh = make(chan *HttpChData, 100)
}
