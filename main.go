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

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}

func main() {
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
}
