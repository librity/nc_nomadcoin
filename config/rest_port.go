package config

const (
	DefaultRestPort = 5001
)

var (
	restPort int
)

func GetRestPort() int {
	if restPort != 0 {
		return restPort
	}

	return DefaultRestPort
}

func SetRestPort(port int) {
	restPort = port
}
