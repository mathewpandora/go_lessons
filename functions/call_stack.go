package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	//LIFO - last in first out
	//call stack
	fn1()
}

func fn1() {
	fn2()
	fmt.Println("fn1")
}

func fn2() {
	fmt.Println("fn2")
	debug.PrintStack() //Печатает какая функция какую вызывала сам стек
}
