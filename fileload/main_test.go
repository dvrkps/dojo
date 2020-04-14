package fileload

import (
	"bufio"
	"fmt"
	"testing"
)

func TestNewReader(t *testing.T) {
	r := newReader(10)

	s := bufio.NewScanner(r)

	i := 0

	for s.Scan() {
		want := fmt.Sprintf(rowFormat, i, i)

		got := s.Text()
		if got != want {
			t.Fatalf("got %q; want %q", got, want)
		}

		i++
	}

	err := s.Err()
	if err != nil {
		t.Errorf("reader: %v", err)
	}
}
