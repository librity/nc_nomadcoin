package p2p

import (
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/librity/nc_nomadcoin/utils"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: checkOrigin,
	}
)

// NOT SAFE AT ALL
func checkOrigin(r *http.Request) bool { return true }

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	seniorConn, err := upgrader.Upgrade(rw, r, nil)
	utils.PanicError(err)

	address := parseAddress(r.RemoteAddr)
	port := utils.GetQuery(r, "thisPort")
	initPeer(address, port, seniorConn)
}

func parseAddress(remoteAddress string) string {
	return strings.Split(remoteAddress, ":")[0]
}
