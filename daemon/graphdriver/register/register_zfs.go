// +build !exclude_graphdriver_zfs,linux !exclude_graphdriver_zfs,freebsd

package register // import "github.com/sdslabs/docker/daemon/graphdriver/register"

import (
	// register the zfs driver
	_ "github.com/sdslabs/docker/daemon/graphdriver/zfs"
)
