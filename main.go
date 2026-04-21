package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("anonymous fn")
	}()

	time.Sleep(time.Second)
}
