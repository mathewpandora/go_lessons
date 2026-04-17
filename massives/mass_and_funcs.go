package main

import "fmt"

func main() {
	//если масив мальенкиц - можно передавать копии (процесс копирования занимает время) если много лучшке апо указателю
	//передача масивов в функцию - по значению (отдается копия)
	nums := [5]int{1, 2, 3, 4, 5}

	modifyArray(nums)

	fmt.Println(nums)

	modifyArrayByPointer(&nums)

	fmt.Println(nums)
}

func modifyArray(arr [5]int) {
	arr[2] = 100
	fmt.Println(arr)
	fmt.Println(&arr)
}

func modifyArrayByPointer(arr *[5]int) {
	arr[2] = 100
	fmt.Println(arr)
}
