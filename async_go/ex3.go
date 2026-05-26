package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100) * time.Now().Second()))
	fmt.Println("Работает менее 3х секунд")
}

func predictableTimeWork() {
	done := make(chan struct{})

	go func() {
		randomTimeWork()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Менее трех секунд")
	case <-time.After(3 * time.Second):
		fmt.Println("Более 3х секунд")
	}
}

func main() {
	predictableTimeWork()
}
