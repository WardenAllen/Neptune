package socket

import (
	"context"
	"errors"
	"net"
	"pluto/common"
	"sync"
	"time"
)

// SocketService struct
type SocketService struct {
	onMessage		func(*Session, *Message)
	onConnect		func(*Session)
	onDisconnect	func(*Session, error)
	onHttpMsg		func(*Session, interface{})
	sessions		*sync.Map
	sessionMap		*sync.Map
	hbInterval		time.Duration
	hbTimeout		time.Duration
	laddr			string
	status			int
	listener		net.Listener
	stopCh			chan error
}

// NewSocketService create a new socket service
func NewSocketService(laddr string) (*SocketService, error) {

	l, err := net.Listen("tcp", laddr)

	if err != nil {
		return nil, err
	}

	s := &SocketService {
		// <SessionID, *Session>
		sessions:   &sync.Map{},
		// <ServerID, *Session>
		sessionMap:	&sync.Map{},
		stopCh:     make(chan error),
		hbInterval: 0 * time.Second,
		hbTimeout:  0 * time.Second,
		laddr:      laddr,
		status:     STInited,
		listener:   l,
	}

	return s, nil
}

// RegMessageHandler register message handler
func (s *SocketService) RegMessageHandler(handler func(*Session, *Message)) {
	s.onMessage = handler
}

// RegConnectHandler register connect handler
func (s *SocketService) RegConnectHandler(handler func(*Session)) {
	s.onConnect = handler
}

// RegDisconnectHandler register disconnect handler
func (s *SocketService) RegDisconnectHandler(handler func(*Session, error)) {
	s.onDisconnect = handler
}

// RegHttpHandler register http handler
func (s *SocketService) RegHttpHandler(handler func(*Session, interface{})) {
	s.onHttpMsg = handler
}

func (s *SocketService) RegServer (sid uint32, session *Session) {

	if _, ok := s.sessionMap.Load(sid); ok {
		// sid already exist.
		return
	}

	s.sessionMap.Store(sid, session)
	session.SetServerID(sid)
}

func (s *SocketService) FindServer(sid uint32) (interface{}, bool) {
	return s.sessionMap.Load(sid)
}

// Start socket service
func (s *SocketService) Start() {

	s.status = STRunning
	ctx, cancel := context.WithCancel(context.Background())

	defer func() {
		s.status = STStop
		cancel()
		_ = s.listener.Close()
	}()

	go s.acceptHandler(ctx)

	for {
		select {

		case <-s.stopCh:
			return

		case p := <-common.HttpCh:

			if s.onHttpMsg != nil {

				if p.Sid != 0 {

					if v,ok := s.sessionMap.Load(p.Sid); ok {
						v.(*Session).GetConn().SendHttpData(p.Data)
					}

				} else {

					// ServerID equals 0 means iterate all sessions.

					f := func(k, v interface{}) bool {

						//这个函数的入参、出参的类型都已经固定，不能修改
						//可以在函数体内编写自己的代码，调用map中的k,v

						s.onHttpMsg(v.(*Session), p.Data)

						return true

					}

					s.sessions.Range(f)
				}

			}

		}
	}
}

func (s *SocketService) acceptHandler(ctx context.Context) {

	for {

		c, err := s.listener.Accept()
		if err != nil {
			s.stopCh <- err
			return
		}

		go s.connectHandler(ctx, c)

	}

}

func (s *SocketService) connectHandler(ctx context.Context, c net.Conn) {

	conn := NewConn(c, s.hbInterval, s.hbTimeout)
	session := NewSession(conn, s)

	s.sessions.Store(session.GetSessionID(), session)

	connctx, cancel := context.WithCancel(ctx)

	defer func() {

		// close operations.
		cancel()
		conn.Close()

		// delete map with this session.
		s.sessions.Delete(session.GetSessionID())
		s.sessionMap.Delete(session.GetServerID())

	}()

	go conn.readCoroutine(connctx)
	go conn.writeCoroutine(connctx)

	if s.onConnect != nil {
		s.onConnect(session)
	}

	for {

		select {

		case err := <-conn.done:

			if s.onDisconnect != nil {
				s.onDisconnect(session, err)
			}
			return

		case msg := <-conn.messageCh:
			if s.onMessage != nil {
				s.onMessage(session, msg)
			}

		case pb := <-conn.httpCh:
			if s.onHttpMsg != nil {
				s.onHttpMsg(session, pb)
			}

		}
	}
}

// GetStatus get socket service status
func (s *SocketService) GetStatus() int {
	return s.status
}

// Stop stop socket service with reason
func (s *SocketService) Stop(reason string) {
	s.stopCh <- errors.New(reason)
}

// SetHeartBeat set heart beat
func (s *SocketService) SetHeartBeat(hbInterval time.Duration, hbTimeout time.Duration) error {
	if s.status == STRunning {
		return errors.New("Can't set heart beat on service running")
	}

	s.hbInterval = hbInterval
	s.hbTimeout = hbTimeout

	return nil
}

// GetConnsCount get connect count
func (s *SocketService) GetConnsCount() int {
	var count int
	s.sessions.Range(func(k, v interface{}) bool {
		count++
		return true
	})
	return count
}

// Unicast Unicast with session ID
func (s *SocketService) Unicast(sid string, msg *Message) {
	v, ok := s.sessions.Load(sid)
	if ok {
		session := v.(*Session)
		err := session.GetConn().SendMessage(msg)
		if err != nil {
			return
		}
	}
}

// Broadcast Broadcast to all connections
func (s *SocketService) Broadcast(msg *Message) {
	s.sessions.Range(func(k, v interface{}) bool {
		s := v.(*Session)
		if err := s.GetConn().SendMessage(msg); err != nil {
			// log.Println(err)
		}
		return true
	})
}
