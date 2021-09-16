package utils

import "strings"

// GetStrChunk safely splits a string by an arbitrary separator
// and returns the chunk at position "index".
func GetStrChunk(str, sep string, index int) string {
	if index < 0 {
		return ""
	}

	splits := strings.Split(str, sep)
	maxIndex := len(splits) - 1
	if index > maxIndex {
		return ""
	}

	return splits[index]
}
