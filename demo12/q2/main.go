package main

import "errors"

type operate func(x, y int) int

func main() {

}

func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}


