// +build !windows

package dockerfile // import "github.com/sdslabs/docker/builder/dockerfile"

func defaultShellForOS(os string) []string {
	return []string{"/bin/sh", "-c"}
}
