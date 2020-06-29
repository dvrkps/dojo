package reader

import (
	"fmt"
	"strings"
)

func read(content string) error {
	r := strings.NewReader(content)

	buf := make([]byte, 3)

	for {
		n, err := r.Read(buf)
		if err != nil {
			return err
		}

		fmt.Printf("read %q, size: %v\n", buf, n)
	}

	return nil
}
