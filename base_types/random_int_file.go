package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	randomNum1 := rand.IntN(100) //[0, 100)
	min := 10
	max := 50
	randomNum2 := rand.IntN(max-min) + min //[10, 50) [0, 40) + 10 (min)

	fmt.Println(randomNum1)
	fmt.Println(randomNum2)
}
