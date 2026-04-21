package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Now it prints")
	time.Sleep(time.Millisecond)
}
