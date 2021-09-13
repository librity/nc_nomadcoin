package p2p

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/librity/nc_nomadcoin/utils"
)

func UpgradePeer(rw http.ResponseWriter, r *http.Request) {
	upgrader, ip, port := buildUpgrader(r)
	juniorConn, err := upgrader.Upgrade(rw, r, nil)
	utils.PanicError(err)

	peer := initPeer(ip, port, juniorConn)
	peer.inbox <- []byte("I am senior, you are junior.")
}

func buildUpgrader(r *http.Request) (*websocket.Upgrader, string, string) {
	ip := parseIP(r.RemoteAddr)
	port := utils.GetQuery(r, "thisPort")
	checkOrigin := func(r *http.Request) bool {
		if ip == "" {
			return false
		}
		if port == "" {
			return false
		}

		return true
	}

	upgrader := &websocket.Upgrader{CheckOrigin: checkOrigin}
	return upgrader, ip, port
}

func parseIP(RemoteAddr string) string {
	return utils.SafeSplit(RemoteAddr, ":", 0)
}
