package main

import (
	"fmt"
	"strings"
	"time"
)

func matchExp(pattern, name string) bool {
	px := 0
	nx := 0
	for px < len(pattern) || nx < len(name) {
		if px < len(pattern) {
			c := pattern[px]
			switch c {
			default: // ordinary character
				if nx < len(name) && name[nx] == c {
					px++
					nx++
					continue
				}
			case '?': // single-character wildcard
				if nx < len(name) {
					px++
					nx++
					continue
				}
			case '*': // zero-or-more-character wildcard
				// Try to match at nx, nx+1, and so on.
				for ; nx <= len(name); nx++ {
					if matchExp(pattern[px+1:], name[nx:]) {
						return true
					}
				}
			}
		}
		// Mismatch.
		return false
	}
	// Matched all of pattern to all of name. Success.
	return true
}

func matchLinear(pattern, name string) bool {
	px := 0
	nx := 0
	nextPx := 0
	nextNx := 0
	for px < len(pattern) || nx < len(name) {
		if px < len(pattern) {
			c := pattern[px]
			switch c {
			default: // ordinary character
				if nx < len(name) && name[nx] == c {
					px++
					nx++
					continue
				}
			case '?': // single-character wildcard
				if nx < len(name) {
					px++
					nx++
					continue
				}
			case '*': // zero-or-more-character wildcard
				// Try to match at nx. If that doesn't work out, restart at nx+1 next.
				nextPx = px
				nextNx = nx + 1
				px++
				continue
			}
		}
		// Mismatch. Maybe restart.
		if 0 < nextNx && nextNx <= len(name) {
			px = nextPx
			nx = nextNx
			continue
		}
		return false
	}
	// Matched all of pattern to all of name. Success.
	return true
}

var tests = []struct {
	pattern string
	name    string
	ok      bool
}{
	{"", "", true},
	{"x", "", false},
	{"", "x", false},
	{"abc", "abc", true},
	{"*", "abc", true},
	{"*c", "abc", true},
	{"*b", "abc", false},
	{"a*", "abc", true},
	{"b*", "abc", false},
	{"a*", "a", true},
	{"*a", "a", true},
	{"a*b*c*d*e*", "axbxcxdxe", true},
	{"a*b*c*d*e*", "axbxcxdxexxx", true},
	{"a*b?c*x", "abxbbxdbxebxczzx", true},
	{"a*b?c*x", "abxbbxdbxebxczzy", false},
	{"a*a*a*a*b", strings.Repeat("a", 100), false},
	{"*x", "xxx", true},
}

func test(desc string, f func(pattern, name string) bool) {
	ok := true
	for _, tt := range tests {
		start := time.Now()
		if f(tt.pattern, tt.name) != tt.ok {
			ok = false
			fmt.Printf("%s(%q, %q) = %v, want %v\n", desc, tt.pattern, tt.name, !tt.ok, tt.ok)
		}
		dt := time.Since(start)
		if dt > 1*time.Millisecond {
			ok = false
			fmt.Printf("%s(%q, %q) took %.3fs - too slow\n", desc, tt.pattern, tt.name, dt.Seconds())
		}
	}
	if ok {
		fmt.Printf("%s ok\n", desc)
	} else {
		fmt.Printf("%s FAIL\n", desc)
	}
}
