// Package main implements experiment with constant errors.
//
// More: http://dave.cheney.net/2016/04/07/constant-errors
package main

import "fmt"

// Error is implementation of error.
type Error string

// Error implements builtin error.
func (e Error) Error() string {
	return string(e)
}

const myError = Error("myerror")

func main() {
	err := magic(-4)

	if err != nil {
		fmt.Println("err != nil", "err =", err)
	}

	if err == myError {
		fmt.Println("err == myError")
	}
}

func magic(i int) error {
	if i < 0 {
		return myError
	}
	return nil
}
