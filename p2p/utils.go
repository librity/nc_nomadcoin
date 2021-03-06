package p2p

import (
	"fmt"

	"github.com/librity/nc_nomadcoin/utils"
)

func buildPeerAdr(ip, port string) string {
	address := fmt.Sprintf("%s:%s", ip, port)

	return address
}

func dismantlePeerAdr(address string) (string, string) {
	ip := utils.GetStrChunk(address, ":", 0)
	port := utils.GetStrChunk(address, ":", 1)

	return ip, port
}
