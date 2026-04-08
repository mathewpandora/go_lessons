package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	res := find_sum(2, 3)
	fmt.Println(res)
	str := "gbfggbgfb вамвамвамавм"
	bytes, runes := getFullLenght(str)
	fmt.Println(bytes, runes)
}

func find_sum(n1, n2 int) int {
	return n1 + n2
}

func getFullLenght(str string) (int, int) {
	return len(str), utf8.RuneCountInString(str)
}
