package main

import "fmt"

func main() {

	nums := [5]int{2, 4, 6, 8, 0}
	fmt.Println(nums[2])
	nums[2] = 34
	fmt.Println(nums[2], nums)
	//len для масива работает нормально
}
