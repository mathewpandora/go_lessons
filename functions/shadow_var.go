package main

import "fmt"

//
//a := 0 так нельзя объявлять вне функции
//damage := 0

func main() {
	a := 5
	if true {
		a := 10 // это новая переменная x, затеняет внешнюю
		fmt.Println(a)
	}
	fmt.Println(a)
}
