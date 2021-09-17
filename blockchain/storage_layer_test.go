package blockchain

type fakeStorageLayer struct {
	fakeLoadChain func() []byte
	fakeLoadBlock func() []byte
}

func (f fakeStorageLayer) LoadChain() []byte {
	return f.fakeLoadChain()
}

func (fakeStorageLayer) SaveChain(chain []byte) {}
func (fakeStorageLayer) ClearChain()            {}

func (f fakeStorageLayer) LoadBlock(hash string) []byte {
	return f.fakeLoadBlock()
}

func (fakeStorageLayer) SaveBlock(hash string, data []byte) {}
func (fakeStorageLayer) ClearBlocks()                       {}
