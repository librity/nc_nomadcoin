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
		return newDifficulty(chain)
	}

	return chain.Dificulty
}

func newDifficulty(chain *blockchain) int {
	currentDificulty := chain.Dificulty
	actualTime := timeSinceLastRecalc(chain)

	if tooEasy(actualTime) {
		return currentDificulty + 1
	}

	if tooHard(actualTime) {
		if currentDificulty > baseDifficulty {
			return currentDificulty - 1
		}
	}

	return currentDificulty
}

func timeSinceLastRecalc(chain *blockchain) int64 {
	blocks := getLastNBlocks(chain, blocksPerRecalc)
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
