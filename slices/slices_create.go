package main

import "fmt"

func main() {
	//у слайсов очень много подводных камней
	arr := []int{} //динамический (работает через масивы)
	//слайс состоит из len cap(сколько элементов не хватат до пересоздания большого массива ) и указател. на масив
	var slice []int
	fmt.Println(slice == nil)
	fmt.Println(arr)

	//make

	slice2 := make([]int, 5, 10) //cap = 10
	fmt.Println(slice2, cap(slice2))

	slice3 := make([]int, 0, 10) //cap = 10 - он пустой но память на 10 элемнтов

	//

	var s []int // zero-value = nil

	fmt.Println(slice3, s)
}
