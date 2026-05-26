package main

//Тут получается паттерн пайплайн - когда данные перетекают от генератора кт генератору
import (
	"fmt"
)

func ex_writer() <-chan int {

	ch := make(chan int)

	go func() {
		defer close(ch)
		for num := range 10 {
			ch <- num
		}

	}()

	return ch
}

func ex_doubler(inputChanel <-chan int) <-chan int {

	ch := make(chan int)

	go func() {
		defer close(ch)
		for v := range inputChanel {
			ch <- v * 2
		}

	}()

	return ch
}

func ex_reader(inputChanel <-chan int) {
	for v := range inputChanel {
		fmt.Println(v)
	}
}

func main() {
	ex_reader(ex_doubler(ex_writer()))
}
