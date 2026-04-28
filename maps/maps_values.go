package main

import "fmt"

func main() {
	//значение в меп может быть любое

	m := make(map[string]map[string]int) //словарь с ключем строкуоой и значением мапой

	m["string1"] = map[string]int{
		"str1": 1,
		"str2": 2,
	}

	m["string2"] = map[string]int{
		"str1": 1,
		"str2": 2,
	}

	fmt.Println(m)
	fmt.Println(m["fruits"] == nil)    //зиро валуе хоть и нил но мапа
	fmt.Println(m["fruits"]["fruits"]) //
	val, ok := m["fruits"]["fruits"]
	fmt.Println(val, ok)

	type my_type map[string]map[string]int //ну типа как с рунами rune int32
	new_map := make(my_type)
	fmt.Println(new_map)
}
