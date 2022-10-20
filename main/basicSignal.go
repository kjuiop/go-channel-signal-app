package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2

	received1 := <-ch
	received2 := <-ch

	ch <- 3
	received3 := <-ch

	fmt.Println(received1)
	fmt.Println(received2)
	fmt.Println(received3)
}
