package utils

import (
	"fmt"
	"testing"
)

func TestGetStrChunk(t *testing.T) {
	goofyStr := "I;;AM;;A;;GOOFY;;;STRINGGGG"

	type testCase struct {
		str      string
		sep      string
		index    int
		expected string
	}

	testCases := []testCase{
		{str: goofyStr, sep: ";;", index: 0, expected: "I"},
		{str: goofyStr, sep: ";;", index: 2, expected: "A"},
		{str: goofyStr, sep: ";;", index: 10, expected: ""},
		{str: goofyStr, sep: ";;", index: -1, expected: ""},
		{str: goofyStr, sep: "B", index: 0, expected: goofyStr},
		{str: goofyStr, sep: "B", index: 1, expected: ""},
	}

	for _, tc := range testCases {
		result := GetStrChunk(tc.str, tc.sep, tc.index)

		if result != tc.expected {
			t.Errorf("Expectd \"%s\", got \"%s\"", tc.expected, result)
		}
	}

}

func ExampleGetStrChunk() {
	test := "I;;AM;;A;;GOOFY;;;STRINGGGG"
	chunk := GetStrChunk(test, ";;", 4)
	fmt.Println(chunk)
	// Output: ;STRINGGGG
}
