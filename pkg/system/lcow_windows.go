package system // import "github.com/sdslabs/docker/pkg/system"

// LCOWSupported returns true if Linux containers on Windows are supported.
func LCOWSupported() bool {
	return lcowSupported
}
