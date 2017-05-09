package main

import (
	"github.com/dvrkps/dojo/initfuncs/aaa"
	"github.com/dvrkps/dojo/initfuncs/zzz"
)

var global string

func init() {
	println("main/main: top init")
}

func main() {
	println("main")
	aaa.Aaa()
	zzz.Zzz()
}

func init() {
	println("main/main: bottom init")
}
