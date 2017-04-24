package main

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
