package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Main goroutine is started ...")

	var wg sync.WaitGroup

	wg.Add(1)
	go task("task1", &wg)

	wg.Add(2)
	go task("task2", &wg)
	go task("task3", &wg)

	fmt.Println("Main goroutine is going to wait to complete sub go routine ...")

	wg.Wait()
	fmt.Println("Main goroutine is done ...")

	var mutex sync.Mutex
	var counter int

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)

	var atomicIntCounter atomic.Int64
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				atomicIntCounter.Add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("atomicIntCounter: ", atomicIntCounter.Load())

}

func task(message string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(message, " : ", i)
	}
}
