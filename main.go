package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Goroutine")
	}()

	fmt.Println("Main")
	time.Sleep(time.Millisecond * 100)
}
