package main

func main() {

	examples := []struct {
		label string
		fn    func()
	}{

		{
			label: "nil interface",
			fn:    ex1,
		},

		{
			label: "interface to value",
			fn:    ex2,
		},
	}

	for _, ex := range examples {
		println(ex.label)
		ex.fn()
		println()
	}
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
