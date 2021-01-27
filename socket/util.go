package socket

import "net"

func Connect(host string) (err error) {

	var (
		tcpAddr		*net.TCPAddr
		conn		*net.TCPConn
		msg			*Message
		data		[]byte
	)

	tcpAddr, err = net.ResolveTCPAddr("tcp", host)
	if err != nil {
		return
	}

	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return
	}

	msg = NewMessage(1, []byte("Hello Zero!"))

	data, err = Encode(msg)

	if err != nil {
		return
	}

	_, _ = conn.Write(data)

	return

}
