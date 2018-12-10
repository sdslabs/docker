package daemon // import "github.com/sdslabs/docker/daemon"

import (
	"github.com/sdslabs/docker/api/types"
	"github.com/sdslabs/docker/pkg/sysinfo"
)

// fillPlatformInfo fills the platform related info.
func (daemon *Daemon) fillPlatformInfo(v *types.Info, sysInfo *sysinfo.SysInfo) {
}

func fillDriverWarnings(v *types.Info) {
}
