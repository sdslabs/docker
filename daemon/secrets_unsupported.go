// +build !linux,!windows

package daemon // import "github.com/sdslabs/docker/daemon"

func secretsSupported() bool {
	return false
}
