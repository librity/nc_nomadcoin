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

func handleAllBlocksReq(message *Msg, p *peer) {
	sendAllBlocksResp(p)
}

func handleAllBlocksResp(message *Msg, p *peer) {
	theirBlocks := []*blockchain.Block{}
	utils.FromJSON(message.Payload, &theirBlocks)

	fmt.Println("🤝 Replacing local blocks with", fmt.Sprintf("%s's", p.address), "blocks.")
	blockchain.Replace(theirBlocks)
}
