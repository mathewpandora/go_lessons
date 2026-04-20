package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	ptr := &numbers[2]
	fmt.Println(ptr, *ptr)
	numbers = append(numbers, 60) //если перебор по капасити то под капотом создатся новый массив
	fmt.Println(ptr, *ptr)

	*ptr = 100 //а он ссылается на элемент в страом масиве поэтмоу numbers не поменяется
	fmt.Println(ptr, &numbers[2])
	fmt.Println(numbers)
}
