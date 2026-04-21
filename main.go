package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}

func send(ch chan<- string) {
	ch <- "send-only-channel"
}

func receive(ch <-chan string) {
	msg := <-ch
	fmt.Println(msg)
}
