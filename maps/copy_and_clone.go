package main

import (
	"fmt"
	"maps"
)

func main() {
	//как скопировать все значения из одногго масива в другой

	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}

	m2 := map[string]int{
		"с": 1,
		"d": 2,
	}

	maps.Copy(m1, m2) //in m1 insert m2 (если есть одинаковые ключи - вставятся ключи из m2)

	fmt.Println(m1)

	m3 := maps.Clone(m2) //клон

	fmt.Println(m3)
}
