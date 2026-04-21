package main

import (
	"fmt"
	"time"
)

func task() {
	time.Sleep(500 * time.Second)
	fmt.Println("done after delay")
}

func main() {
	go task()
	time.Sleep(time.Second)
}
