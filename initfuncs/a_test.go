package main

import "testing"

func init() {
	println("test main/a: top init")
}

func TestA(t *testing.T) {
	println("TestA")
}

func init() {
	println("test main/a: bottom init")
}
