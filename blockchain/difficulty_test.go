package blockchain

import (
	"testing"

	"github.com/librity/nc_nomadcoin/utils"
)

func TestGetDifficulty(t *testing.T) {

	t.Run("Should increase difficulty if too easy", func(t *testing.T) {
		blockControl := 0
		fakeBlocks := []*Block{
			{Hash: "5", PreviousHash: "4", Timestamp: now() + 6},
			{Hash: "4", PreviousHash: "3", Timestamp: now() + 5},
			{Hash: "3", PreviousHash: "2", Timestamp: now() + 4},
			{Hash: "2", PreviousHash: "1", Timestamp: now() + 1},
			{Hash: "1", PreviousHash: "", Timestamp: now()},
		}

		storage = fakeStorageLayer{
			fakeLoadBlock: func() []byte {
				defer func() { blockControl++ }()

				return utils.ToGob(fakeBlocks[blockControl])
			},
		}
		chain := &blockchain{LastHash: "5", Dificulty: baseDifficulty}

		type testCase struct {
			height   int
			expected int
		}
		testCases := []testCase{
			{height: 0, expected: baseDifficulty},
			{height: 2, expected: baseDifficulty},
			{height: 5, expected: baseDifficulty + 1},
			{height: 6, expected: baseDifficulty + 1},
		}

		for _, tc := range testCases {
			chain.Height = tc.height
			result := getDifficulty(chain)
			chain.Dificulty = result

			if result != tc.expected {
				t.Log(tc)
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		}
	})

	t.Run("Should decrease difficulty if too hard", func(t *testing.T) {
		blockControl := 0
		fakeBlocks := []*Block{
			{Hash: "11", PreviousHash: "10", Timestamp: now() + 350*60},
			{Hash: "10", PreviousHash: "9", Timestamp: now() + 300*60},
			{Hash: "9", PreviousHash: "8", Timestamp: now() + 250*60},
			{Hash: "8", PreviousHash: "7", Timestamp: now() + 200*60},
			{Hash: "7", PreviousHash: "6", Timestamp: now() + 150*60},
			{Hash: "6", PreviousHash: "5", Timestamp: now() + 100*60},
			{Hash: "5", PreviousHash: "4", Timestamp: now() + 6},
			{Hash: "4", PreviousHash: "3", Timestamp: now() + 5},
			{Hash: "3", PreviousHash: "2", Timestamp: now() + 4},
			{Hash: "2", PreviousHash: "1", Timestamp: now() + 1},
			{Hash: "1", PreviousHash: "", Timestamp: now()},
		}

		storage = fakeStorageLayer{
			fakeLoadBlock: func() []byte {
				defer func() { blockControl++ }()

				return utils.ToGob(fakeBlocks[blockControl])
			},
		}
		chain := &blockchain{LastHash: "11", Dificulty: baseDifficulty + 1}

		type testCase struct {
			height   int
			expected int
		}
		testCases := []testCase{
			{height: 9, expected: baseDifficulty + 1},
			{height: 10, expected: baseDifficulty},
			{height: 11, expected: baseDifficulty},
		}

		for _, tc := range testCases {
			chain.Height = tc.height
			result := getDifficulty(chain)
			chain.Dificulty = result

			if result != tc.expected {
				t.Log(tc)
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		}
	})
}

func Test(t *testing.T) {

	t.Run("", func(t *testing.T) {

	})

}
