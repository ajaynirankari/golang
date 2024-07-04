package main

import "fmt"

func main() {

	ch := make(chan int, 1)

	go sendOnlyGoRoutine(ch, 45)
	go receiveOnlyGoRoutine(ch)
	go bidirectionalGoRoutine(ch)

	// Wait for goroutines to finish
	fmt.Scanln()

}
func sendOnlyGoRoutine(ch chan<- int, data int) {
	// ch is a send-only channel
	fmt.Println("Writing data to channel")
	ch <- data
	fmt.Println("Data written to channel")

	//value := <-ch   // Error: invalid operation: <-ch (receive from send-only type chan<- int)
	//fmt.Println("writeDataToChannel can read data from channel", value)
}
func receiveOnlyGoRoutine(ch <-chan int) {
	// ch is a read-only channel
	fmt.Println("Reading data from channel")
	println(<-ch)
	fmt.Println("Data read from channel")

	//ch <- 100    // Error: invalid operation: ch <- 100 (send to receive-only type <-chan int)
	//fmt.Println("readDataFromChannel can write data to channel")
}
func bidirectionalGoRoutine(ch chan int) {
	// ch is a bidirectional channel
	fmt.Println("Writing data to channel")
	ch <- 100
	fmt.Println("Data written to channel")

	fmt.Println("Reading data from channel")
	println(<-ch)
	fmt.Println("Data read from channel")
}
