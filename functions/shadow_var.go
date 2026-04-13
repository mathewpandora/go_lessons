package main

import "fmt"

func main() {
	a := 5
	if true {
		a := 10
		fmt.Println(a)
	}
	fmt.Println(a)
}
