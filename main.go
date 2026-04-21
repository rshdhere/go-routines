package main

import (
	"fmt"
	"time"
)

func heavy() {
	for i := 0; i < 1e7; i++ {
	}
	fmt.Println("heavy done")
}

func main() {
	go heavy()
	fmt.Println("main continues")
	time.Sleep(time.Second)
}
