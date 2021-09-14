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

	theirHeight := theirLastBlock.Height
	ourHeight := blockchain.GetBC().Height
	if theirHeight == ourHeight {
		fmt.Println("On same height as", p.address)
		return
	}

	if theirHeight > ourHeight {
		fmt.Println(p.address, "is ahead by", theirHeight-ourHeight, "blocks")
		sendAllBlocksReq(p)
		return
	}

	fmt.Println(p.address, "is behind by", ourHeight-theirHeight, "blocks")
	sendLastBlock(p)
}

func handleAllBlocksReq(message *Msg, p *peer) {
	sendAllBlocksResp(p)
}

func handleAllBlocksResp(message *Msg, p *peer) {
	theirBlocks := &[]blockchain.Block{}
	utils.FromJSON(message.Payload, theirBlocks)
	fmt.Println("Their blocks:", theirBlocks)

	// blockchain.Restore(theirBlocks)
}
