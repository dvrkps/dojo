package aaa

func init() {
	println("aaa/main: top init")
}

func Aaa() {
	println("aaa")
}

func init() {
	println("aaa/main: bottom init")
}
