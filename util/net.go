package util

func GetServerInfo(sid uint32) (channel uint32, stype uint32, id uint32) {

	channel = sid >> 16
	stype = (sid >> 8) & 0xFF
	id = sid & 0xFF

	return

}
