package rest

import "strings"

func cleanPort() string {
	return strings.Trim(port, ":")
}
