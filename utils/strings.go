package utils

import "strings"

// GetChunk safely splits a string by an arbitrary separator
// and returns the chunk at position "index".
func GetChunk(str, sep string, index int) string {
	splits := strings.Split(str, sep)
	maxIndex := len(splits) - 1
	if index > maxIndex {
		return ""
	}

	return splits[index]
}
