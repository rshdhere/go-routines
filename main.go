package main

import (
	"fmt"
	"time"
)

func run() {
	go fmt.Println("inside fn")
}

func main() {
	run()
	time.Sleep(time.Second)
}
