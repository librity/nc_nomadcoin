package blockchain

import "fmt"

func AddPeerBlock(peerBlock *Block) {
	getBC().addPeerBlock(peerBlock)
}

func (b *blockchain) addPeerBlock(peerBlock *Block) {
	// TODO: Verify block

	peerBlock.save()
	b.reference(peerBlock)

	getMP().removeConfirmedTxs(peerBlock)
}

func Replace(blocks []*Block) {
	getBC().replace(blocks)
}

func (b *blockchain) replace(blocks []*Block) {
	b.reset()

	for _, block := range blocks {
		block.save()
	}

	lastBlock := blocks[0]
	b.reference(lastBlock)

	fmt.Println("⛓️  Blockchain replaced succesfully.")
}
