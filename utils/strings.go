package utils

import "strings"

func SafeSplit(str, sep string, index int) string {
	splits := strings.Split(str, sep)
	maxIndex := len(splits) - 1
	if index > maxIndex {
		return ""
	}

	return splits[index]
}
