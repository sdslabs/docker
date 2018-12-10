package daemon // import "github.com/sdslabs/docker/daemon"

import (
	"testing"

	"github.com/sdslabs/docker/api/types"
	"github.com/sdslabs/docker/dockerversion"
	"gotest.tools/assert"
)

func TestFillLicense(t *testing.T) {
	v := &types.Info{}
	d := &Daemon{
		root: "/var/lib/docker/",
	}
	d.fillLicense(v)
	assert.Assert(t, v.ProductLicense == dockerversion.DefaultProductLicense)
}
