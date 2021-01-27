package socket

import (
	"fmt"
)

// Message struct
type Message struct {
	msgLen		uint32
	msgID		uint16
	tranID		uint32
	data		[]byte
}

// NewMessage create a new message
func NewMessage(msgID uint16, data []byte) *Message {

	msg := &Message{
		msgLen:		uint32(len(data)) + 6,
		msgID:		msgID,
		tranID:		uint32(len(data)),
		data:		data,
	}

	return msg
}

// GetData get message data
func (msg *Message) GetData() []byte {
	return msg.data
}

// GetID get message ID
func (msg *Message) GetID() uint16 {
	return msg.msgID
}

func (msg *Message) String() string {
	return fmt.Sprintf("Size=%d ID=%d DataLen=%d",
		msg.msgLen, msg.GetID(), len(msg.GetData()))
}
