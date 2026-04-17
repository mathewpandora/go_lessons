package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{5, 6, 7}
	fmt.Println(arr1 == arr2) //сравниваются поэлементно
	fmt.Println(arr1 == arr3)
}
