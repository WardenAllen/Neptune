package util

import (
	"fmt"
	"math/big"
	"net"
)

func InetNtoA(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip), byte(ip>>8), byte(ip>>16), byte(ip>>24))
}

func InetAtoN(ip string) uint32 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return uint32(ret.Uint64())
}

