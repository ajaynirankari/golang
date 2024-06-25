package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	go sayGoodbye()

	go func() {
		fmt.Println("Anonymous function executed")
	}()

	go func(n int) {
		fmt.Println("Anonymous function executed with argument:", n)
	}(4)

	ch := make(chan string)

	go sendMessages(ch)
	go receiveMessages(ch)
	go receiveMessages(ch)

	
	ch1 := make(chan string, 5)
	close(ch1)

	time.Sleep(100 * time.Second)

	fmt.Println("Main function executed")

}
func sendMessages(ch chan string) {
	ch <- "Hello"
	ch <- "World"
	close(ch)
}

func receiveMessages(ch chan string) {
	msg := <-ch
	fmt.Println(msg)
}

func sayHello() string {
	fmt.Println("Hello")
	return "Hello"
}

func sayGoodbye() string {
	fmt.Println("Goodbye")
	return "Goodbye"
}
