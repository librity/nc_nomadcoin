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
	theirBlock := &blockchain.Block{}
	utils.FromJSON(message.Payload, theirBlock)

	handleBlockDifferences(p, theirBlock)
}

func handleBlockDifferences(p *peer, theirBlock *blockchain.Block) {
	ourBlock := blockchain.GetLastBlock()

	if theirBlock.Height == ourBlock.Height {
		handleSameHeight(p, theirBlock, ourBlock)
		return
	}

	if theirBlock.Height > ourBlock.Height {
		fmt.Println("🤝", p.address, "is ahead by", theirBlock.Height-ourBlock.Height, "blocks.")
		sendAllBlocksReq(p)
		return
	}

	fmt.Println("🤝", p.address, "is behind by", ourBlock.Height-theirBlock.Height, "blocks.")
	sendLastBlock(p)
}

func handleSameHeight(p *peer, theirBlock, ourBlock *blockchain.Block) {
	fmt.Println("🤝 On same height as", p.address)

	if theirBlock.Hash != ourBlock.Hash {
		fmt.Println("🤝 Requesting blocks from", p.address, "due to hash differences.")
		sendAllBlocksReq(p)
		return
	}

	fmt.Println("🤝 On same hash as", p.address)
}

func handleAllBlocksReq(message *Msg, p *peer) {
	sendAllBlocksResp(p)
}

func handleAllBlocksResp(message *Msg, p *peer) {
	theirBlocks := &[]*blockchain.Block{}
	utils.FromJSON(message.Payload, theirBlocks)

	fmt.Println("🤝 Their blocks:", *theirBlocks)

	// blockchain.Restore(theirBlocks)
}
