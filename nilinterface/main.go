package main

func main() {

	examples := []struct {
		label string
		fn    func()
	}{

		{
			label: "nil interface",
			fn:    toNil,
		},

		{
			label: "interface to value",
			fn:    toValue,
		},

		{
			label: "interface to nil pointer",
			fn:    toNilPointer,
		},

		{
			label: "interface to pointer",
			fn:    toPointer,
		},
	}

	for _, ex := range examples {
		println(ex.label)
		ex.fn()
		println()
	}
}

func isNil(i interface{}) {
	println("is nil:", i == nil)
}

func toNil() {
	var i interface{} = nil
	isNil(i)   // true
	println(i) // (0x0, 0x0)
}

func toValue() {
	var v int = 10
	var i interface{} = v
	isNil(i) // false
	println("value address:", &v)
	println(i)
}

func toNilPointer() {
	var p *int = nil
	var i interface{} = p
	isNil(i) // false
	println("nil ptr:", p)
	println(i)
}

func toPointer() {
	var v int = 10
	var p *int = &v
	var i interface{} = p
	isNil(i) // false
	println("pointer:", p)
	println(i)
}
