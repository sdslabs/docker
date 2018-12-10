// +build !linux

package archive // import "github.com/sdslabs/docker/pkg/archive"

func getWhiteoutConverter(format WhiteoutFormat) tarWhiteoutConverter {
	return nil
}
