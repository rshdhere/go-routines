package main

import (
	"fmt"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Printf("%s %d\n", msg, i)
		time.Sleep(time.Second)
	}
}

func main() {
	boring("boring!")
}
