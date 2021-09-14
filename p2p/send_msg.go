package p2p

import "github.com/librity/nc_nomadcoin/blockchain"

func sendLastBlock(p *peer) {
	lastBlock := blockchain.GetLastBlock()
	message := makeMsgJSON(MsgLastBlock, lastBlock)

	p.inbox <- message
}

func sendAllBlocksReq(p *peer) {

}

func sendAllBlocksResp(p *peer) {

}
