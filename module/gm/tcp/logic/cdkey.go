package logic

import (
	"bytes"
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"github.com/golang/protobuf/proto"
	"pluto/database"
	"pluto/log"
	"pluto/module/gm/config"
	Proto "pluto/protobuf"
	"pluto/socket"
	"pluto/util"
	"strconv"
	"strings"
)

const (
	protoTCPCdkeyReq		= "TCPCdkeyReq"
	rc4Key					= "WardenAllen"
)

type CDKeyType int32
const (
	CDKeyTypeNone			CDKeyType = 0
	CDKeyTypeOnce			CDKeyType = 1
	CDKeyTypeCommon			CDKeyType = 2
)

type CDKeyState int32
const (
	CDKeyStateNone			CDKeyState = 0
	CDKeyStateValid			CDKeyState = 1
	CDKeyStateUsed			CDKeyState = 2
	CDKeyStateInvalid		CDKeyState = 3
)

func OnTCPCdkeyReq(s *socket.Session, msg *socket.Message) bool {

	var (
		err			error
		req			*Proto.CdkeyReq
		ack			*Proto.CdkeyAck
		strMd5		string
		strRC4		string
		strSign		string
		rc4Byte		[]byte
		dstByte		[]byte
		ktype		uint8
		keyid		uint32
		channel		uint8
		awardId		uint16
		binBuf		*bytes.Buffer
		rows		*sql.Rows
		rs			sql.Result
		affectNum	int64
		errCode		int32
		stateStr	string
		state		int

		s1			string
		s2			string
		s3			string
		s4			string
		s5			string
		s6			string
		s7			string

		cdkeyAward	*config.CDKeyAward
		awards		[]*Proto.AwardItem
		data		[]byte

		keyGenInfo	*config.CDKeyGeneralInfo
	)

	errCode = 0

	req = &Proto.CdkeyReq{}
	err = proto.Unmarshal(msg.GetData(), req)

	if err != nil {
		log.Proto(protoTCPCdkeyReq, false, "Parse failed.")
		errCode = 1
		goto FINISH
	}

	// if it's a general cdkey.
	keyGenInfo = config.GetCdkeyGenInfo(req.CDKey)
	if keyGenInfo != nil {

		awardId = uint16(keyGenInfo.AwardID)
		ktype = uint8(CDKeyTypeCommon)

		goto FINISH
	}

	if len(req.CDKey) != 20 {
		log.Proto(protoTCPCdkeyReq, false, "Invalid CDKey length.")
		errCode = 1
		goto FINISH
	}

	s1 = req.CDKey[:1]
	s2 = req.CDKey[1:2]
	s3 = req.CDKey[2:3]
	s4 = req.CDKey[3:17]
	s5 = req.CDKey[17:18]
	s6 = req.CDKey[18:19]
	s7 = req.CDKey[19:]

	strMd5 = s1 + s3 + s5 + s7
	strRC4 = s2 + s4 + s6

	strSign = util.EncryptMd5(strRC4)
	if len(strSign) <= 4 {
		log.Proto(protoTCPCdkeyReq, false, "EncryptMd5 failed.")
		errCode = 1
		goto FINISH
	}

	strSign = strings.ToUpper(strSign[:4])
	if strMd5 != strSign {
		log.Proto(protoTCPCdkeyReq, false, "Sign failed.")
		errCode = 1
		goto FINISH
	}

	rc4Byte, err = hex.DecodeString(strings.ToLower(strRC4))
	if err != nil {
		log.Proto(protoTCPCdkeyReq, false, "Hex decode failed.")
		errCode = 1
		goto FINISH
	}

	dstByte, err = util.DecryptRC4Byte(rc4Byte, rc4Key)
	if err != nil {
		log.Proto(protoTCPCdkeyReq, false, "RC4 decode failed.")
		errCode = 1
		goto FINISH
	}

	if len(dstByte) != 8 {
		log.Proto(protoTCPCdkeyReq, false, "Invalid byte length.")
		errCode = 1
		goto FINISH
	}

	binBuf = bytes.NewBuffer(dstByte[:1])
	err = binary.Read(binBuf, binary.BigEndian, &ktype)
	if err != nil {
		log.Proto(protoTCPCdkeyReq, false, "invalid ktype.")
		errCode = 1
		goto FINISH
	}

	binBuf = bytes.NewBuffer(dstByte[1:5])
	err = binary.Read(binBuf, binary.BigEndian, &keyid)
	if err != nil {
		log.Proto(protoTCPCdkeyReq, false, "invalid keyid.")
		errCode = 1
		goto FINISH
	}

	binBuf = bytes.NewBuffer(dstByte[5:6])
	err = binary.Read(binBuf, binary.BigEndian, &channel)
	if err != nil {
		log.Proto(protoTCPCdkeyReq, false, "invalid channel.")
		errCode = 1
		goto FINISH
	}

	binBuf = bytes.NewBuffer(dstByte[6:8])
	err = binary.Read(binBuf, binary.BigEndian, &awardId)
	if err != nil {
		log.Proto(protoTCPCdkeyReq, false, "invalid awardId.")
		errCode = 1
		goto FINISH
	}

	// TODO : check keyid.

	database.LockMysql()
	defer database.UnlockMysql()

	// check cdkey's state

	rows, err = database.MySQLConns[0].Query("SELECT state FROM cdkey WHERE keyid = ?", keyid)

	if err != nil {
		log.Proto(protoTCPCdkeyReq, false, "Mysql query failed.")
		errCode = 1
		goto FINISH
	}

	if rows == nil {
		/* invalid cdkey */
		errCode = 1
		goto FINISH
	}

	for rows.Next() {
		err = rows.Scan(&stateStr)
		if err != nil {
			log.Proto(protoTCPCdkeyReq, false, "Mysql row invalid.")
			errCode = 1
			goto FINISH
		}
	}

	state, err = strconv.Atoi(stateStr)
	if err != nil {
		log.Proto(protoTCPCdkeyReq, false, "Mysql row state invalid.")
		errCode = 1
		goto FINISH
	}

	_ = rows.Close()

	if state != 1 {

		/* key is already used */
		errCode = 2
		goto FINISH
	}

	/* CDKeyTypeOnce need change state */

	if CDKeyType(ktype) == CDKeyTypeOnce {

		rs, err = database.MySQLConns[0].Exec(
			"UPDATE cdkey SET state=? WHERE keyid = ?", CDKeyStateUsed, keyid)

		if err != nil {
			errCode = 1
			goto FINISH
		}

		if rs == nil {
			errCode = 1
			goto FINISH
		}

		affectNum, err = rs.RowsAffected()

		if affectNum == 0 {
			errCode = 1
			goto FINISH
		}

	}

FINISH:

	if errCode == 0 {

		/* verify success, send awards */

		cdkeyAward = config.GetCdkeyAward(int(awardId))

		if cdkeyAward != nil {

			if cdkeyAward.Type != int(ktype) {

				errCode = 1

			} else {

				awards = cdkeyAward.Award

			}

		} else {

			errCode = 1

		}

	}

	ack = &Proto.CdkeyAck{
		Error:		errCode,
		CDKey: 		req.CDKey,
		UserID:		req.UserID,
		KeyType:	int32(ktype),
		Awards:		awards,
	}

	data, err = proto.Marshal(ack)

	if err != nil {
		log.Proto(protoTCPCdkeyReq, false, "Marshal failed.")
		return false
	}

	msg = socket.NewMessage(uint16(Proto.PLUTO_ID_CDKeyAckID), data)

	err = s.GetConn().SendMessage(msg)

	if err != nil {
		log.Proto(protoTCPCdkeyReq, false, "Send failed.")
		return false
	}

	log.Proto(protoTCPCdkeyReq, true, "uid %d cdkey %s err %d.", req.UserID, req.CDKey, errCode)

	return true
}