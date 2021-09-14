package utils

import "fmt"

func BuildPort(portNum int) string {
	port := fmt.Sprintf(":%d", portNum)

	return port
}
