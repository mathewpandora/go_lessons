package main

import "maps"

func main() {
	m := map[string][]int{
		"a": {5, 2},
		"b": {20, 11},
		"c": {1, 3, 9},
		"d": {6, 111, 5},
	}

	//Удалим все пары ключ-значение если есть число 5 слайсе
	// лучше всего создать новый меп но можно и так:

	maps.DeleteFunc(m, func(key string, value []int) bool {
		for _, v := range value {
			if v == 5 {
				return true
			}
		}
		return false

		//slices.Contains(value, 5)
	})

}
