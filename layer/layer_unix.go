// +build linux freebsd darwin openbsd

package layer // import "github.com/sdslabs/docker/layer"

import "github.com/sdslabs/docker/pkg/stringid"

func (ls *layerStore) mountID(name string) string {
	return stringid.GenerateRandomID()
}
