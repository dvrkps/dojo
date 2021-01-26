package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println("vim-go")
}

func IsIt(t *testing.T, fail bool) {
	t.Helper()
	if fail {
		t.Error("ups")
	}
}
