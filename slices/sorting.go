package main

import (
	"fmt"
	"slices"
)

func main() {
	inst := []int{9, 12, 45, 54, 23, 12}
	slices.Sort(inst)
	fmt.Println(inst)
	nums := []int{4, 7, 9, 2, 3, 3, 1}
	slices.SortFunc(nums, func(a, b int) int { return a - b })
	slices.SortStableFunc()
}

func comporator(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}

}
