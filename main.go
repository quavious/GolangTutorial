package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	name := "lee"
	totalLength, _ := lenAndUpper("nico")
	fmt.Println(name)
	fmt.Println(multiply(2, 2))
	fmt.Println(totalLength)
	repeatMe("nico", "lin", "super", "ill")
}
