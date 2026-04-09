package main

import "fmt"

func main() {
	fmt.Println(fn())
}

func fn() (i int) {
	i = 5
	defer func() {
		i = 6
	}()
	return 7
}
