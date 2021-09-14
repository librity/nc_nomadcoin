package config

import "sync"

const (
	DefaultExplorerPort = 4000
)

var (
	explorerPort int
	explorerOnce sync.Once
)

func GetExplorerPort() int {
	if explorerPort == 0 {
		return DefaultExplorerPort
	}

	return explorerPort
}

func SetExplorerPort(port int) {
	explorerOnce.Do(func() {
		explorerPort = port
	})
}
