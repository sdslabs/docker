package daemon // import "github.com/sdslabs/docker/daemon"

import (
	"testing"

	containertypes "github.com/sdslabs/docker/api/types/container"
	"github.com/sdslabs/docker/container"
	"github.com/sdslabs/docker/daemon/config"
	"github.com/sdslabs/docker/daemon/exec"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestGetInspectData(t *testing.T) {
	c := &container.Container{
		ID:           "inspect-me",
		HostConfig:   &containertypes.HostConfig{},
		State:        container.NewState(),
		ExecCommands: exec.NewStore(),
	}

	d := &Daemon{
		linkIndex:   newLinkIndex(),
		configStore: &config.Config{},
	}

	_, err := d.getInspectData(c)
	assert.Check(t, is.ErrorContains(err, ""))

	c.Dead = true
	_, err = d.getInspectData(c)
	assert.Check(t, err)
}
