package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(n int) {
			time.Sleep(time.Duration(n) * 200 * time.Millisecond)
			fmt.Println(n)
		}(i)
	}
	time.Sleep(2 * time.Second)
}
