package config

import (
	"fmt"
	"sync"
)

const (
	DefaultRestPort = 5001
)

var (
	restPort int
	restOnce sync.Once
)

func GetRestPort() int {
	if restPort == 0 {
		return DefaultRestPort
	}

	return restPort
}

func GetRestPortStr() string {
	portStr := fmt.Sprint(GetRestPort())

	return portStr
}

func SetRestPort(port int) {
	restOnce.Do(func() {
		restPort = port

		SetDBName()
	})
}
