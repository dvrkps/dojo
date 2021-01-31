package storage

import (
	"fmt"
	"path"
)

func UserPath(root string, user string) string {
	p := fmt.Sprintf("%s/pills/%s", root, user)

	return path.Clean(p)
}
