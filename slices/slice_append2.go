package main

import (
	"fmt"
)

func main() {
	slcie := []int{0, 0, 0}
	fmt.Println(slcie)
	slcie = append(slcie, 4) //кладем в конец
	fmt.Println(slcie)
	slcie = append([]int{6}, slcie...) //кладем в 0 нашщ слайс (короче вставили ноль в начало) начало
	fmt.Println(slcie)
	index := len(slcie) / 2
	value := 777
	before := slcie[:index]
	after := append([]int{value}, slcie[index:]...) //то что вставляем плюс то чт после
	fmt.Println(before, after)
	slcie = append(before, after...)
	fmt.Println(slcie)

}
