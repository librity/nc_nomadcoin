package explorer

import "time"

func increment(number int) int {
	return number + 1
}

func add(a, b int) int {
	return a + b
}

func unixToHuman(unix int64) string {
	return time.Unix(unix, 0).Format(time.UnixDate)
}

func blockURL(hash string) string {
	return baseURL + "/blocks/" + hash
}
