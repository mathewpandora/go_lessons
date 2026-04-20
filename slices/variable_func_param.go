package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 7}
	fmt.Println(sum(s...)) //spread... оператор - распаковка слайса
}

func sum(values ...int) int { //values - это слайс как раз
	sum := 0
	for _, v := range values {
		fmt.Println(v)
		sum += v
	}
	return sum

}
