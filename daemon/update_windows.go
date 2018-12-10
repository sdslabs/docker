package daemon // import "github.com/sdslabs/docker/daemon"

import (
	"github.com/sdslabs/docker/api/types/container"
	"github.com/sdslabs/docker/libcontainerd"
)

func toContainerdResources(resources container.Resources) *libcontainerd.Resources {
	// We don't support update, so do nothing
	return nil
}
