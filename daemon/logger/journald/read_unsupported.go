// +build !linux !cgo static_build !journald

package journald // import "github.com/sdslabs/docker/daemon/logger/journald"

func (s *journald) Close() error {
	return nil
}
