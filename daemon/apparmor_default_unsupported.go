// +build !linux

package daemon // import "github.com/sdslabs/docker/daemon"

func ensureDefaultAppArmorProfile() error {
	return nil
}
