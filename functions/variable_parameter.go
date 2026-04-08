package main

import "fmt"

func main() {
	sum1(12, 5, 6, 7)
	sum2(1, 2)
}

// параметры с переменным числом агрументов:
func sum1(values ...int) { // они хранятся слайсом
	fmt.Println(values)
}

func sum2(a1 int, a2 int, values ...int) {
	fmt.Println(a1, a2)
}
