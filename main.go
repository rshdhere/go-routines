package main

import (
	"fmt"
	"time"
)

func work(id int) {
	fmt.Println(id)
}

func main() {
	for i := 1; i <= 3; i++ {
		go work(i)
	}

	time.Sleep(time.Second)
}
