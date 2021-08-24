package libcore

import (
	"github.com/sagernet/libping"
	"os"
)

func init() {
	initLog()
}

func Setenv(key, value string) error {
	return os.Setenv(key, value)
}

func Unsetenv(key string) error {
	return os.Unsetenv(key)
}

var ipv6Mode int

func SetIPv6Mode(mode int) {
	ipv6Mode = mode
}

func IcmpPing(address string, timeout int) (int, error) {
	return libping.IcmpPing(address, timeout)
}
