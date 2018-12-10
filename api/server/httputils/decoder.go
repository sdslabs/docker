package httputils // import "github.com/sdslabs/docker/api/server/httputils"

import (
	"io"

	"github.com/sdslabs/docker/api/types/container"
	"github.com/sdslabs/docker/api/types/network"
)

// ContainerDecoder specifies how
// to translate an io.Reader into
// container configuration.
type ContainerDecoder interface {
	DecodeConfig(src io.Reader) (*container.Config, *container.HostConfig, *network.NetworkingConfig, error)
	DecodeHostConfig(src io.Reader) (*container.HostConfig, error)
}
