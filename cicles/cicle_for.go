package main

import "fmt"

func main() {
	// DRY - dont repeat yourself - не надо дублировать код
	for i := 0; i < 5; i++ { //i++ срабатывает только после тела
		fmt.Println("привет!")
	}

	i1 := 0
	for ; i1 < 5; i1++ { //i++ срабатывает только после тела
		fmt.Println("привет!")
	}
	fmt.Println(i1)

	i2 := 0
	for i2 < 5 { //while аналог
		fmt.Println("dfsv")
		i2++
	}

	for a, b := 0, 0; a < 10 && b > 3; a, b = a+1, b-1 {
		fmt.Printf("a=%d and b=%d", a, b)
	}

	for el := range 10 {
		fmt.Println(el)
	}

	for range 10 {
		fmt.Println("10 times")
	}
}
