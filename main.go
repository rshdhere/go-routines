package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"zeta",
		"pi",
		"epsilon",
		"theta",
	}

	wg.Add(len(words))

	for index, value := range words {
		go printSomething(fmt.Sprintf("%d : %s", index, value), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("This is ths second thing to be printed", &wg)
}
