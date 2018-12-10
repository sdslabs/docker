// +build !linux

package vfs // import "github.com/sdslabs/docker/daemon/graphdriver/vfs"

import "github.com/sdslabs/docker/pkg/chrootarchive"

func dirCopy(srcDir, dstDir string) error {
	return chrootarchive.NewArchiver(nil).CopyWithTar(srcDir, dstDir)
}
