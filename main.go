package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("outter")
		go func() {
			fmt.Println("inner")
		}()
	}()

	time.Sleep(time.Second)
}
