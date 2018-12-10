package register // import "github.com/sdslabs/docker/daemon/graphdriver/register"

import (
	// register the windows graph drivers
	_ "github.com/sdslabs/docker/daemon/graphdriver/lcow"
	_ "github.com/sdslabs/docker/daemon/graphdriver/windows"
)
