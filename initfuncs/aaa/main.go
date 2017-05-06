package aaa

func init() {
	println("aaa/main: top init")
}

func Aaa() {
	println("aaa/main: main")
}

func init() {
	println("aaa/main: bottom init")
}
