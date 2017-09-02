package main

import "bytes"

func better(parts []string) []string {
	found := false
	var buffer bytes.Buffer
	var newParts []string

	for _, val := range parts {
		if found {
			newParts = append(newParts, val)
			continue
		}
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
	}

	return newParts
}
