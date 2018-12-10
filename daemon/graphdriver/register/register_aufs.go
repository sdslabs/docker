// +build !exclude_graphdriver_aufs,linux

package register // import "github.com/sdslabs/docker/daemon/graphdriver/register"

import (
	// register the aufs graphdriver
	_ "github.com/sdslabs/docker/daemon/graphdriver/aufs"
)
