// +build !exclude_graphdriver_overlay,linux

package register // import "github.com/sdslabs/docker/daemon/graphdriver/register"

import (
	// register the overlay graphdriver
	_ "github.com/sdslabs/docker/daemon/graphdriver/overlay"
)
