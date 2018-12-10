package plugins // import "github.com/sdslabs/docker/pkg/plugins"

import (
	"os"
	"path/filepath"
)

var specsPaths = []string{filepath.Join(os.Getenv("programdata"), "docker", "plugins")}
