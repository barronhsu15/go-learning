package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return a / b, nil
}

func main() {
	result, err := divide(10, 0)

	if err == nil {
		fmt.Println("Result: ", result)
	} else {
		fmt.Println("error: ", err)
	}
}
