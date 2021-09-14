package p2p

import (
	"fmt"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/utils"
)

func handleMsg(message *Msg, p *peer) {
	switch message.Kind {
	case MsgLastBlock:
		handleLastBlock(message, p)
	case MsgAllBlocksReq:
		handleAllBlocksReq(message, p)
	case MsgAllBlocksResp:
		handleAllBlocksResp(message, p)
	}
}

func handleLastBlock(message *Msg, p *peer) {
	theirLastBlock := &blockchain.Block{}
	utils.FromJSON(message.Payload, theirLastBlock)
	fmt.Println(theirLastBlock)
}

func handleAllBlocksReq(message *Msg, p *peer) {
}

func handleAllBlocksResp(message *Msg, p *peer) {
}
