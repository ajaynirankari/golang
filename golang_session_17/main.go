package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		time.Sleep(1 * time.Second)
	}()

	go func() {
		ch2 <- 2
		time.Sleep(2 * time.Second)
	}()

	go sendData(ch2)

	select {
	case x, ok := <-ch1:
		if !ok {
			fmt.Println("Channel closed")
			return
		}
		fmt.Println("received value from ch1: ", x)
	case x := <-ch2:
		fmt.Println("received value from ch2: ", x)
	default:
		fmt.Println("None of the channels ready to communicate")
	}

	time.Sleep(5 * time.Second)

	select {
	case ch1 <- 100:
		fmt.Println("send value on ch1: ")
	case ch2 <- 200:
		fmt.Println("send value on ch2: ")
	default:
		fmt.Println("None of the channels ready to communicate")
		time.Sleep(1 * time.Second)
	}

	time.Sleep(5 * time.Second)

	select {
	case ch1 <- 100:
		fmt.Println("send value on ch1: ")
	case x := <-ch2:
		fmt.Println("receive value on ch2: ", x)
	default:
		fmt.Println("None of the channels ready to communicate")
		time.Sleep(1 * time.Second)
	}

}

func sendData(ch chan int) {
	time.Sleep(2 * time.Second)
	ch <- 2
}
