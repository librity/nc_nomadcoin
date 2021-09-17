package db

type DBLayer struct{}

func (DBLayer) LoadChain() []byte {
	return loadChain()
}

func (DBLayer) SaveChain(chain []byte) {
	saveChain(chain)
}

func (DBLayer) ClearChain() {
	clearChain()
}

func (DBLayer) LoadBlock(hash string) []byte {
	return loadBlock(hash)
}

func (DBLayer) SaveBlock(hash string, data []byte) {
	saveBlock(hash, data)
}

func (DBLayer) ClearBlocks() {
	clearBlocks()
}
