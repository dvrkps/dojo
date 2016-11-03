package main

func main() {

	examples := []struct {
		label string
		fn    func()
	}{

		{
			label: "nil interface",
			fn:    nilInterface,
		},

		{
			label: "interface to value",
			fn:    interfaceToValue,
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

func nilInterface() {
	var i interface{} = nil
	println("is nil:", i == nil) // true
	println(i)                   // (0x0, 0x0)
}

func interfaceToValue() {
	var v int = 10
	var i interface{} = v
	println("is nil:", i == nil) // false
	println("value address:", &v)
	println(i)
}

func toNilPointer() {
	var p *int = nil
	var i interface{} = p
	println("is nil:", i == nil) // false
	println("nil ptr:", p)
	println(i)
}

func toPointer() {
	var v int = 10
	var p *int = &v
	var i interface{} = p
	println("is nil:", i == nil) // false
	println("pointer:", p)
	println(i)
}
