package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Herllo"
	}()

	go func() {
		ch2 <- "World"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("received", msg1)
	case msg2 := <-ch2:
		fmt.Println("received", msg2)
	}

}
