package main

import "testing"

func init() {
	println("test main/z: top init")
}

func TestZ(t *testing.T) {
	println("TestZ")
}

func init() {
	println("test main/z: bottom init")
}
