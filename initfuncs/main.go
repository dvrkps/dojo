package main

import (
	"github.com/dvrkps/dojo/initfuncs/aaa"
	"github.com/dvrkps/dojo/initfuncs/zzz"
)

var global = "default"

func init() {
	println("main/main: top init")
	setGlobal("mit")
}

func main() {
	println("main")
	setGlobal("mm")
	aaa.Aaa()
	zzz.Zzz()
}

func init() {
	println("main/main: bottom init")
	setGlobal("mib")
}

func setGlobal(s string) {
	old := global
	global = s
	println("new: ", global, " old: ", old)
}
