// +build !linux,!windows

package sysinfo // import "github.com/sdslabs/docker/pkg/sysinfo"

import (
	"runtime"
)

// NumCPU returns the number of CPUs
func NumCPU() int {
	return runtime.NumCPU()
}
