package explorer

import (
	"fmt"
	"time"
)

func debug(i interface{}) string {
	log := fmt.Sprintf("DEBUG: %v\n", i)
	fmt.Print(log)

	return log
}

func increment(number int) int {
	return number + 1
}

func add(a, b int) int {
	return a + b
}

func unixToHuman(unix int64) string {
	return time.Unix(unix, 0).Format(time.UnixDate)
}

func homeURL() string {
	return baseURL
}

func blockURL(hash string) string {
	return baseURL + "/blocks/" + hash
}

func txURL(hash string) string {
	return baseURL + "/transactions/" + hash
}
