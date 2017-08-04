package main

func main() {}

type T struct {
	I int64   `json:"i,omitempty"`
	F float64 `json:"f,omitempty"`
	B bool    `json:"b,omitempy"`
	S []byte  `json:"s,omitempty"`
}
