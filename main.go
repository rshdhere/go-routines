package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	go printSomething("This is ths first thing to be printed")

	time.Sleep(1 * time.Second)

	printSomething("This is ths second thing to be printed")
}
