package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}

	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}

	fmt.Println(maps.Equal(m1, m2))

	//задача сравнить по ключам и сравнить суммы за7чнеимй слайсов

	m3 := map[string][]int{
		"a": {5, 2},
		"b": {20, 11},
	}

	m4 := map[string][]int{
		"a": {5, 1, 1},
		"b": {11, 2, 18},
	}

	results := maps.EqualFunc(m3, m4, func(v1, v2 []int) bool { //принимает два значения по одному ключу
		sum1 := 0

		for _, v := range v1 {
			sum1 += v
		}

		sum2 := 0

		for _, v := range v2 {
			sum2 += v
		}

		if sum1 == sum2 {
			return true
		}

		return false
	})

	fmt.Println(results)
}
