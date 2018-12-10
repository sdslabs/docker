// +build !linux,!windows

package daemon // import "github.com/sdslabs/docker/daemon"

func configsSupported() bool {
	return false
}
