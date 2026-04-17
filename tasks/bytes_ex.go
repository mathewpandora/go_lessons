package main

import "fmt"

func main() {
	num := 5
	num = num << 2
	fmt.Println(num)
	num = num >> 1
	fmt.Println(num)
	num = num & 3
	fmt.Println(num)
	num = num | 3
	fmt.Println(num)
	num = num ^ 2
	num = ^num
	fmt.Println(num)
}
