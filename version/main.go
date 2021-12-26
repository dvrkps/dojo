package main

import (
	"runtime/debug"
)

func main() {
	println(version())
}

func version() string {
	const v = "v0.1.0"

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return v
	}

	s := bi.Settings

	var commit = ""
	for i := range s {
		if s[i].Key == "vcs.revision" {
			commit = s[i].Value
			break
		}
	}

	if commit == "" {
		return v
	}

	return v + "-" + commit
}
