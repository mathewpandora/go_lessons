package main

import "fmt"

func main() {

	nums := [5]int{12, 45, 76, 76, 56}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for index, value := range nums {
		fmt.Println(index, value)
	}
}
