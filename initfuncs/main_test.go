package main

import "testing"

func init() {
	println("test main/main: top init")
}

func TestM(t *testing.T) {
	println("TestM")
}

func init() {
	println("test main/main: bottom init")
}
