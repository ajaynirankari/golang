package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go producer(ch, &wg)

	wg.Add(1)
	go consumer(ch, &wg)

	//time.Sleep(6 * time.Second)
	fmt.Println("main function going to wait ...")
	wg.Wait()
	fmt.Println("main function exit")

}

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("producer produce value : ", i)
		time.Sleep(2 * time.Second)
	}
	close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("consumer consume value : ", val)
	}
}
