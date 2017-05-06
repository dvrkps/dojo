package main

import "github.com/dvrkps/dojo/initfuncs/aaa"

func init() {
	println("main/main: top init")
}

func main() {
	println("main")
	aaa.Aaa()
}

func init() {
	println("main/main: bottom init")
}
