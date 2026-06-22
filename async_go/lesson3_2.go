package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := range 3 {
			fmt.Println(i, "Запись")
			ch <- i //Она встанет и будет висель вечно в заблокированнои состоянии ее надо  закрыть 
		}
	}()
	// Дедлок — это ситуация, когда все горутины заблокированы навечно и ни одна не может проснуться. Здесь же main не заблокирован (он просто спит, но это конечное ожидание)
	go func() {
		for v := range ch {
			fmt.Println(v, "Чтение")
			return
		}
	}()

	time.Sleep(10 * time.Second)
}
