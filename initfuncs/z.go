package main

func init() {
	println("main/z: top init")
	setGlobal("zit")
}

func init() {
	println("main/z: bottom init")
	setGlobal("zib")
}
