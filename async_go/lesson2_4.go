package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := range 100000 {
			ch <- i
		}
		wg.Done()
	}()

	go func() {
		for i := range 100000 {
			ch <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	go func() {
		for v := range ch {
			fmt.Println("v=", v, "worker1")
		}
	}()

	for v := range ch {
		fmt.Println("v=", v, "worker2")

	}
}
