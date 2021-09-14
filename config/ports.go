package config

const (
	DefaultExplorerPort = 4000
	DefaultRestPort     = 5001
)

var (
	explorerPort int
	restPort     int
)

func GetExplorerPort() int {
	if explorerPort != 0 {
		return explorerPort
	}

	return DefaultExplorerPort
}

func GetRestPort() int {
	if restPort != 0 {
		return restPort
	}

	return DefaultRestPort
}

func SetExplorerPort(port int) {
	explorerPort = port
}

func SetRestPort(port int) {
	restPort = port
}
