package utils

import "fmt"

// BuildPort buils a port string from a port number.
func BuildPort(portNum int) string {
	port := fmt.Sprintf(":%d", portNum)

	return port
}
