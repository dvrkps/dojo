package main

import "bytes"

func main() {}

func short() []string {
	return []string{".hi", "\"My", "name", "is", "Omar\"", "\"123\""}
}

func long() []string {
	return []string{"\".hi", "I'm", "the", "real", "Slim", "Shady", "\"My", "name", "is", "Omar\"", "hello", "world", "\"123\"", "a"}
}

// Original finds a "string which spans multiple spaces" in a split message.
// Then takes that and replaces the Quote string with a single string value of the quote contents.
func original(parts []string) []string {
	// https://play.golang.org/p/908AF_h0q-
	found := false
	var buffer bytes.Buffer
	var newParts []string

	for _, val := range parts {
		if !found {
			if val[0] == '"' {
				if val[len(val)-1] == '"' {
					found = true
					val = val[:len(val)-1]
					buffer.WriteString(val[1:])
					newParts = append(newParts, buffer.String())
				} else {
					buffer.WriteString(val[1:])
				}
			} else if buffer.Len() != 0 {
				if val[len(val)-1] == '"' {
					found = true
					val = val[:len(val)-1]
					buffer.WriteString(" " + val)
					newParts = append(newParts, buffer.String())
				} else {
					buffer.WriteString(" " + val)
				}
			} else {
				newParts = append(newParts, val)
			}
		} else {
			newParts = append(newParts, val)
		}
	}

	return newParts
}
