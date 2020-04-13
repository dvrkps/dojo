package fileload

import (
	"fmt"
)

// String returns string representation.
func (r Row) String() string {
	s := fmt.Sprintf("id: %02d name: %q", r.ID, r.Name)
	return s
}
