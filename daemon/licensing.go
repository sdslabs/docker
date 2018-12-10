package daemon // import "github.com/sdslabs/docker/daemon"

import (
	"github.com/sdslabs/docker/api/types"
	"github.com/sdslabs/docker/dockerversion"
)

func (daemon *Daemon) fillLicense(v *types.Info) {
	v.ProductLicense = dockerversion.DefaultProductLicense
}
