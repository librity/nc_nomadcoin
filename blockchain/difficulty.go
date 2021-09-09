package blockchain

const (
	baseDifficulty  = 2
	blocksPerRecalc = 5
	minutesPerBlock = 2
	tolerance       = 2
	expectedTime    = blocksPerRecalc * minutesPerBlock
)

func getDifficulty(chain *blockchain) int {
	if chain.Height == 0 {
		return baseDifficulty
	}

	shouldRecalc := chain.Height%blocksPerRecalc == 0
	if shouldRecalc {
		return newDifficulty()
	}

	return chain.CurrentDificulty
}

func newDifficulty() int {
	currentDificulty := GetBC().CurrentDificulty
	actualTime := timeSinceLastRecalc()

	if tooEasy(actualTime) {
		return currentDificulty + 1
	}

	if tooHard(actualTime) {
		return currentDificulty - 1
	}

	return currentDificulty
}

func timeSinceLastRecalc() int64 {
	blocks := LastNBlocks(blocksPerRecalc)
	lastBlock := blocks[0]
	lastRecalcBlock := blocks[blocksPerRecalc-1]
	actualTime := lastBlock.Timestamp - lastRecalcBlock.Timestamp
	actualTime = actualTime / 60

	return actualTime
}

func tooEasy(actualTime int64) bool {
	return actualTime <= (expectedTime - tolerance)
}

func tooHard(actualTime int64) bool {
	return actualTime >= (expectedTime + tolerance)
}
