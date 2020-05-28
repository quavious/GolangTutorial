package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I am done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	total, upper := lenAndUpper("nico")
	fmt.Println(total, upper)
}
