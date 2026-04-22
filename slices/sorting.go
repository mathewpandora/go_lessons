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
}
