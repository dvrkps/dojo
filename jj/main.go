package main

func main() {}

type One struct {
	I int64   `json:"i,omitempty"`
	F float64 `json:"f,omitempty"`
	B bool    `json:"b,omitempy"`
	S string  `json:"s,omitempty"`
}
