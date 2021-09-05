package blockchain

const (
	baseDifficulty  = 2
	blocksPerRecalc = 5
	minutesPerBlock = 2
	tolerance       = 2
	expectedTime    = blocksPerRecalc * minutesPerBlock
)

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return baseDifficulty
	}

	if b.Height%blocksPerRecalc == 0 {
		return b.recalculateDifficulty()
	}

	return b.CurrentDificulty
}

func (b *blockchain) recalculateDifficulty() int {
	actualTime := b.timeSinceLastRecalc()

	if tooEasy(actualTime) {
		return b.CurrentDificulty + 1
	}

	if tooHard(actualTime) {
		return b.CurrentDificulty - 1
	}

	return b.CurrentDificulty
}

func (b *blockchain) timeSinceLastRecalc() int {
	blocks := b.LastNBlocks(blocksPerRecalc)
	lastBlock := blocks[0]
	lastRecalcBlock := blocks[blocksPerRecalc-1]
	actualTime := lastBlock.Timestamp - lastRecalcBlock.Timestamp
	actualTime = actualTime / 60

	return actualTime
}

func tooEasy(actualTime int) bool {
	return actualTime <= (expectedTime - tolerance)
}

func tooHard(actualTime int) bool {
	return actualTime >= (expectedTime + tolerance)
}
