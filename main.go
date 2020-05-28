package main

import "fmt"

func main() {
	name := [5]string{"nico", "lynn", "dal", "kim"}
	name[4] = "shin"
	names := []string{"a", "b", "c", "d"}
	names = append(names, "e")
	fmt.Println(name, names)
}
