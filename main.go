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

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 1)
	c <- person + " is sexy"
}

func main() {
	c := make(chan string)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}
