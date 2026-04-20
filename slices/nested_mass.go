package main

import "fmt"

func main() {
	//вложенные масивы
	//arr := [3]int{1, 2, 3}
	arr := [3][3]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println(arr, arr[1][0])

	arr[2] = [3]int{9, 9, 9}

	fmt.Println(arr, arr[1][0])
	//вложенные слайсы

	slice := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{1, 2, 3},
		{},
	}

	fmt.Println(slice)

	slice[2] = append(slice[2], 200, 300, 400)

	for _, v := range slice {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
}
