package main

import "fmt"

func main() {
	// рекурсия
	fn4(12)
}

func fn3() { //бексонечная рекурсия
	fmt.Println() //в кол стеке это разные фукнции с разным указателем (это так?) каждая
	fn3()
}

func fn4(count int) { //бексонечная рекурсия
	if count < 0 {
		fmt.Println("Првет", count)
		return //стопер
	}
	fmt.Println("Првет", count) //в кол стеке это разные фукнции с разным указателем (это так?) каждая
	fn4(count - 1)
}

func factorial(n int) int { //шаг рекурсии - когда снова вызываем функцию
	if n == 0 { //базовый случай - когда стопер срабатывает
		return 1
	}
	return n * factorial(n-1) //шаги : 5-1 4-1 3-1 2-1 1 ...
}
