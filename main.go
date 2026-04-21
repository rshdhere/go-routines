package main

import (
	"fmt"
	"time"
)

func sayHellow() {
	fmt.Println("hello from the go routine")
}

func main() {
	go sayHellow()
	time.Sleep(time.Second)
}
