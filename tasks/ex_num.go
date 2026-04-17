package main

import (
	"fmt"
	"math/rand/v2"
)

func rollDice(n int) {
	try_count := 0
	for {
		res1, res2 := rand.IntN(12-2)+2+1, rand.IntN(12-2)+2+1
		if res1 == res2 {
			fmt.Printf("Выпало %d и %d, в сумме %d, на это потребовалось %d броска.", res1, res2, res1+res2, try_count)
			break
		} else {
			fmt.Printf("Выпало %d и %d, в сумме %d, бросаем еще раз.\n", res1, res2, res1+res2)
			try_count++
		}
	}
}
