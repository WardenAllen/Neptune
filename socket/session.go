package socket

import (
	uuid "github.com/satori/go.uuid"
)

// Session struct
type Session struct {
	sessionID	string
	serverID	uint32
	conn		*Conn
	service		*SocketService
	settings	map[string]interface{}
}

// NewSession create a new session
func NewSession(conn *Conn, ss *SocketService) *Session {
	id:= uuid.NewV4()
	session := &Session{
		sessionID:	id.String(),
		serverID:	0,
		conn:		conn,
		service:	ss,
		settings:	make(map[string]interface{}),
	}

	return session
}

func (s *Session) GetService() *SocketService {
	return s.service
}

func (s *Session) GetSessionID() string {
	return s.sessionID
}

func (s *Session) SetServerID(sid uint32) {
	s.serverID = sid
}

func (s *Session) GetServerID() uint32 {
	return s.serverID
}

// GetConn get zero.Conn pointer
func (s *Session) GetConn() *Conn {
	return s.conn
}

// SetConn set a zero.Conn to session
func (s *Session) SetConn(conn *Conn) {
	s.conn = conn
}

// GetSetting get setting
func (s *Session) GetSetting(key string) interface{} {

	if v, ok := s.settings[key]; ok {
		return v
	}

	return nil
}

// SetSetting set setting
func (s *Session) SetSetting(key string, value interface{}) {
	s.settings[key] = value
}
