package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	counter := 20
	wg.Add(20)
	for i := 1; i <= counter; i++ {

		go func() {
			defer wg.Done()
			fmt.Println(i * i)
		}()

	}

	wg.Wait()
}
