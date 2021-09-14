package config

import "sync"

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

func SetRestPort(port int) {
	restOnce.Do(func() {
		restPort = port

		SetDBName()
	})
}
