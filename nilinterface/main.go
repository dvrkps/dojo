package main

func main() {
	ex1()

}

func ex1() {
	var i interface{} = nil
	println("nil interface is nil?:", i == nil) // true
	println(i)                                  // (0x0, 0x0)

}
