package main

import (
	"fmt"
	"time"
)

func _count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, ": in ", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go _count("nico")
	go _count("lynn")
	time.Sleep(time.Second * 5)
}
