package config

const (
	DefaultExplorerPort = 4000
)

var (
	explorerPort int
)

func GetExplorerPort() int {
	if explorerPort != 0 {
		return explorerPort
	}

	return DefaultExplorerPort
}

func SetExplorerPort(port int) {
	explorerPort = port
}
