package main

import "fmt"

func main() {
	numbers := []int{1, 3, 3, 3, 5, 5, 5, 5, 5}

	m := map[int]int{}

	for i := 0; i < len(numbers); i++ {
		m[numbers[i]]++
	}

	for key, value := range m {
		fmt.Printf("The %d has %d times\n", key, value)
	}
}
