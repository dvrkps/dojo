package main

import (
	"bytes"
	"strings"
)

func better(in []string) []string {
	found := false
	var buffer bytes.Buffer
	out := make([]string, 0, len(in))

	const q string = `"`

	var size, lastPos int
	//var first, last byte
	for _, row := range in {
		size = len(row)
		lastPos = size - 1
		//first = row[0]
		//last = row[lastPos]
		if found {
			out = append(out, row)
			continue
		}
		if strings.HasPrefix(row, q) {
			if strings.HasSuffix(row, q) {
				found = true
				row = row[:lastPos]
				buffer.WriteString(row[1:])
				out = append(out, buffer.String())
			} else {
				buffer.WriteString(row[1:])
			}
		} else if buffer.Len() != 0 {
			if strings.HasSuffix(row, q) {
				found = true
				row = row[:lastPos]
				buffer.WriteString(" " + row)
				out = append(out, buffer.String())
			} else {
				buffer.WriteString(" " + row)
			}
		} else {
			out = append(out, row)
		}
	}

	return out
}
