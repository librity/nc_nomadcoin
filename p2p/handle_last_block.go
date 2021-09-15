package p2p

import (
	"fmt"

	"github.com/librity/nc_nomadcoin/blockchain"
	"github.com/librity/nc_nomadcoin/utils"
)

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
