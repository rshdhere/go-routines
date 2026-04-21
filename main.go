package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Println("running...")
			time.Sleep(200 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
}
