package main

import (
	"fmt"
	"strings"
)

func reverseString(text string) string {
	textSplit := strings.Split(text, "")
	result := ""

	for i := len(textSplit) - 1; i >= 0; i-- {
		result += textSplit[i]
	}

	return result
}

func main() {
	fmt.Println(reverseString("Hello, World!"))
}
