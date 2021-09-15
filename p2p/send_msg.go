package p2p

import "github.com/librity/nc_nomadcoin/blockchain"

func sendLastBlock(p *peer) {
	lastBlock := blockchain.GetLastBlock()
	message := makeMsgJSON(MsgLastBlock, lastBlock)

	p.inbox <- message
}

func sendAllBlocksReq(p *peer) {
	message := makeMsgJSON(MsgAllBlocksReq, nil)

	p.inbox <- message
}

func sendAllBlocksResp(p *peer) {
	blocks := blockchain.GetBlocks()
	message := makeMsgJSON(MsgAllBlocksResp, blocks)

	p.inbox <- message
}

func sendMinedBlock(p *peer, minedBlock *blockchain.Block) {
	message := makeMsgJSON(MsgMinedBlock, minedBlock)

	p.inbox <- message
}
