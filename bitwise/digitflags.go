package main

import "fmt"

type optionValue int

func (ov optionValue) String() string {
	v := "optionValue can:\n"
	if (ov & optionRead) > 0 {
		v += "can read\n"
	}
	if (ov & optionWrite) > 0 {
		v += "can write\n"
	}
	if (ov & optionExecute) > 0 {
		v += "can execute\n"
	}
	return v
}

const (
	optionRead    optionValue = 1 << iota // 1
	optionWrite                           // 2
	optionExecute                         // 4
)

func newOptionValue(r, w, e bool) optionValue {
	var ov optionValue
	if r {
		ov |= optionRead
	}
	if w {
		ov |= optionWrite
	}
	if e {
		ov |= optionExecute
	}
	return ov
}

func digitFlags() {

	o := newOptionValue(true, false, true)
	fmt.Println(o)

	o2 := optionRead | optionWrite | optionExecute
	fmt.Println(o2)

}
