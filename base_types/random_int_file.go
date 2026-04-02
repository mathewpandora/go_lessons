package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	randomNum1 := rand.IntN(100) //[0, 100)
	min := 10
	max := 50
	randomNum2 := rand.IntN(max - min) //[0, 100)

	fmt.Println(randomNum1)
	fmt.Println(randomNum2)
}
