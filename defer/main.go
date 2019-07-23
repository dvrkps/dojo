package main

func main() {
	foo(43)
}

func foo(f int) {
	defer func(o, f int) {
		println("defer", o, f)
	}(one(), f+10)
	println("foo")
	f++
	println("end", f)
}

func one() int {
	println("one")
	return 1
}
