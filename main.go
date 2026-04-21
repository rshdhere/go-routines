package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("A%d\n", i)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("B%d\n", i)
		}
	}()

	time.Sleep(time.Second)
}
