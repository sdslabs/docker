// +build !linux,!freebsd,!windows

package daemon // import "github.com/sdslabs/docker/daemon"
import "github.com/sdslabs/docker/daemon/config"

const platformSupported = false

func setupResolvConf(config *config.Config) {
}
