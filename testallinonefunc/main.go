package main

import (
	"errors"
	"fmt"
)

func main() {
	i, err := inc(1)
	fmt.Println(i, err)
}

const incMax = 99

func inc(i int) (int, error) {

	if i < 0 {
		return 0, errors.New("negative value")
	}

	if i > incMax {
		return 0, fmt.Errorf("value %d > %d", i, incMax)
	}

	i++

	return i, nil
}
