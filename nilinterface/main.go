package main

func main() {

	ex1()

	ex2()

}

func ex1() {
	var i interface{} = nil
	println("ex1")
	println("nil interface is nil?:", i == nil) // true
	println(i)                                  // (0x0, 0x0)
	println()

}

func ex2() {
	var v int = 10
	var i interface{} = v
	println("ex2")
	println("interface-to-value is nil?:", i == nil) // false
	println("value address:", &v)
	println(i)
	println()
}
