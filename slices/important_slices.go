package main

import "fmt"

func main() {
	//нюансы
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s[0]) //это очень быстрая операция
	s = append(s, 6)
}

func fn(a [5]int) {
	fmt.Println(a)
}
