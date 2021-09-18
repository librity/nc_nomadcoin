package utils

import (
	"math/big"
	"testing"
)

const (
	intsHex = "224c4ab3ef55080f2fd26c9268e5d134f7cafc0ab44badc2b721d4aaea88bf6ba86225643d2410d7309f5f3fe323fe317b16eb314356f5672ef4b9c9044a75a5"
)

var (
	expectedA, _ = new(big.Int).SetString("15513432809018947966887860098053332473297935411847581563380302711326998314859", 10)
	expectedB, _ = new(big.Int).SetString("76161967641385419271517418501539349830171482872916448622942251509704085370277", 10)
)

func TestBigIntsFromHex(t *testing.T) {

	t.Run("Should return the expected big ints", func(t *testing.T) {
		a, b := BigIntsFromHex(intsHex)

		FailIfDifferent(t, 0, a.Cmp(expectedA))
		FailIfDifferent(t, 0, b.Cmp(expectedB))

	})

	t.Run("Should panic when hex isn't valid", func(t *testing.T) {
		badHexes := []string{
			"224c4ab3ef55080f2fd26c9268e5d134f7cafc0ab44badc2b721d4aaea88bf6ba86225643d2410d7309f5f3fe323fe317b16eb314356f5672ef4b9c9044a75a",
			"224c4ab3ef55080f2fd26c9268e5d134f7cafc0ab44badc2b721d4aaea88bf6ba86225643d2410d7309f5f3fe323fe317b16eb314356f5672ef4b9c9044a75a55",
			"224c4ab3ef55",
			"-o2orfoksd-f 0-=1221",
			"",
		}

		for _, bh := range badHexes {
			test := func() {
				BigIntsFromHex(bh)
			}

			ShouldPanic(t, test)
		}
	})

}

func TestBytesToBigInt(t *testing.T) {
	type testCase struct {
		bytes    []byte
		expected *big.Int
	}

	expectedA, _ := new(big.Int).SetString("7984307400472207895118760978068116400306460148948252534383314142104575909921", 10)
	expectedB, _ := new(big.Int).SetString("62500125198449156401328606281050875724028775603087117224202336383151198639538", 10)
	testCases := []testCase{
		{
			bytes:    []byte{17, 166, 245, 53, 48, 39, 184, 175, 95, 27, 103, 222, 37, 67, 203, 55, 50, 149, 83, 239, 78, 181, 149, 247, 178, 46, 182, 20, 212, 153, 160, 33},
			expected: expectedA,
		},

		{
			bytes:    []byte{138, 45, 209, 56, 12, 87, 210, 38, 124, 74, 212, 242, 42, 245, 132, 179, 33, 251, 191, 193, 223, 135, 5, 172, 135, 130, 28, 174, 169, 135, 21, 178},
			expected: expectedB,
		},
	}

	t.Run("Should return the expected big ints", func(t *testing.T) {
		for _, tc := range testCases {
			result, err := BytesToBigInt(tc.bytes)

			FailIfDifferent(t, nil, err)
			FailIfDifferent(t, 0, result.Cmp(tc.expected))

		}
	})

	t.Run("Should validate byte array length", func(t *testing.T) {
		badBytes := [][]byte{
			{17, 166, 245, 53, 48, 39, 184, 175, 95, 27, 103, 222, 37, 67, 203, 55, 50, 149, 83, 239, 78, 181, 149, 247, 178, 46, 182, 20, 212, 153, 160},
			{138, 45, 209, 56, 12, 87, 210, 38, 124, 74, 212, 242, 42, 245, 132, 179, 33, 251, 191, 193, 223, 135, 5, 172, 135, 130, 28, 174, 169, 135, 21, 178, 42},
		}

		for _, bb := range badBytes {
			_, err := BytesToBigInt(bb)

			FailIfDifferent(t, ErrBigIntBadBytes, err)

		}
	})
}
