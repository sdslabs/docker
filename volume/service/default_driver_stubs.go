// +build !linux,!windows

package service // import "github.com/sdslabs/docker/volume/service"

import (
	"github.com/sdslabs/docker/pkg/idtools"
	"github.com/sdslabs/docker/volume/drivers"
)

func setupDefaultDriver(_ *drivers.Store, _ string, _ idtools.Identity) error { return nil }
