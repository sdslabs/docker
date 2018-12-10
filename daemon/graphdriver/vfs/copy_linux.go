package vfs // import "github.com/sdslabs/docker/daemon/graphdriver/vfs"

import "github.com/sdslabs/docker/daemon/graphdriver/copy"

func dirCopy(srcDir, dstDir string) error {
	return copy.DirCopy(srcDir, dstDir, copy.Content, false)
}
