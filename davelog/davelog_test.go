// Package davelog implements a minimal logging package.
// Inspired by article of Dave Cheney.
// https://dave.cheney.net/2015/11/05/lets-talk-about-logging
package davelog

import "testing"

func TestNew(t *testing.T) {
	l := New(nil, false)
	if l == nil {
		t.Error("New( ... ) == nil")
	}
}
