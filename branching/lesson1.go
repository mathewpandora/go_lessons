package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Printf("Введите число: ")
	fmt.Scanln(&a)

	if a > 5 {
		fmt.Println("Число а Больше 5")
	} else if a == 5 {
		fmt.Println("Число а равно 5")
	} else {
		fmt.Println("Число а меньше 5")
	}

	fmt.Println("Теперь с условиями:")

	var b int
	fmt.Printf("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Printf("Введите второе число: ")
	fmt.Scanln(&b)

	if a == 5 && b != 7 {
		fmt.Println("a == 5 и b != 7")
	} else if a > 100 || b < 0 {
		fmt.Println("a > 100 || b < 0")
	} else {
		fmt.Println("Все остальное")
	}

	fmt.Println("Теперь с вложенными:")

	var name string
	fmt.Printf("Введите имя:")
	fmt.Scanln(&name)

	if name != "Andrey" {
		fmt.Println("Не андрей")
		if name != "Vlad" {
			fmt.Println("И не Влад")
		}
	}

}
