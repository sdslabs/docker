// +build !linux

package daemon // import "github.com/sdslabs/docker/daemon"

// ModifyRootKeyLimit is a noop on unsupported platforms.
func ModifyRootKeyLimit() error {
	return nil
}
