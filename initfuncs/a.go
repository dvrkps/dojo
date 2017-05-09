package main

func init() {
	println("main/a: top init")
	setGlobal("ait")
}

func init() {
	println("main/a: bottom init")
	setGlobal("aib")
}
