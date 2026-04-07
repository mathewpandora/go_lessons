package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	//перемнную лучше создавать внутри блока елс
	if num := rand.IntN(100); num > 5 { //[0, n)  Если передать n <= 0 программа вызовет panic
		fmt.Println("Число а Больше 5")
	} else if num == 5 {
		fmt.Println("Число а равно 5")
	} else {
		fmt.Println("Число а меньше 5")
	}

	//if err := doSomething(); err != nil {
	//	fmt.Printf("Произошла ошибка %v", err)
	//}

}
