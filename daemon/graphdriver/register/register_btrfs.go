// +build !exclude_graphdriver_btrfs,linux

package register // import "github.com/sdslabs/docker/daemon/graphdriver/register"

import (
	// register the btrfs graphdriver
	_ "github.com/sdslabs/docker/daemon/graphdriver/btrfs"
)
