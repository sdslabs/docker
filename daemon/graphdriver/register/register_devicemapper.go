// +build !exclude_graphdriver_devicemapper,!static_build,linux

package register // import "github.com/sdslabs/docker/daemon/graphdriver/register"

import (
	// register the devmapper graphdriver
	_ "github.com/sdslabs/docker/daemon/graphdriver/devmapper"
)
