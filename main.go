package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("one")
	go fmt.Println("two")
	go fmt.Println("third")

	time.Sleep(time.Second)
}
