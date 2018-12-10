package daemon // import "github.com/sdslabs/docker/daemon"

import (
	"github.com/sdslabs/docker/container"
	"github.com/sdslabs/docker/pkg/archive"
)

func (daemon *Daemon) tarCopyOptions(container *container.Container, noOverwriteDirNonDir bool) (*archive.TarOptions, error) {
	return daemon.defaultTarCopyOptions(noOverwriteDirNonDir), nil
}
