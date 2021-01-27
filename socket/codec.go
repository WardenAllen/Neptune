package socket

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const (
	Offset = 6
)

// Encode from Message to []byte
func Encode(msg *Message) ([]byte, error) {
	buffer := new(bytes.Buffer)

	err := binary.Write(buffer, binary.BigEndian, msg.msgLen)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buffer, binary.BigEndian, msg.msgID)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buffer, binary.BigEndian, msg.tranID)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buffer, binary.BigEndian, msg.data)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// Decode from []byte to Message
func Decode(buf []byte) (*Message, error) {

	bufReader := bytes.NewReader(buf)

	var (
		msgLen		uint32
		msgID		uint16
		tranID		uint32
		dataLen		int
		err			error
	)

	// 读取消息长度
	err = binary.Read(bufReader, binary.BigEndian, &msgLen)
	if err != nil {
		return nil, err
	}

	// 读取消息ID
	err = binary.Read(bufReader, binary.BigEndian, &msgID)
	if err != nil {
		return nil, err
	}

	// 读取传输ID
	err = binary.Read(bufReader, binary.BigEndian, &tranID)
	if err != nil {
		return nil, err
	}

	// 读取数据
	dataLen = int(msgLen) - Offset
	if dataLen < 0 {
		return nil, errors.New("invalid data length")
	}
	data := make([]byte, dataLen)
	err = binary.Read(bufReader, binary.BigEndian, &data)
	if err != nil {
		return nil, err
	}

	msg := &Message{
		msgLen:  msgLen,
		msgID:   msgID,
		tranID:  tranID,
		data:    data,
	}

	return msg, nil
}
