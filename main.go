package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second)
		fmt.Println("1 sec")
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2 sec")
	}()

	time.Sleep(3 * time.Second)
}
