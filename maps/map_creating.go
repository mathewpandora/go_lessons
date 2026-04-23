package main

import "fmt"

func main() {
	var m map[string]int       //ключ - стринг значение инт
	fmt.Println(m == nil)      //стоит нил
	m2 := make(map[string]int) //можно указать капасити
	m3 := map[string]int{
		"apple":  5,
		"banana": 2,
	}
	fmt.Println(m2, m3) //

	m2["kikik"] = 23 //

	fmt.Println(m["str"]) //если такого ключа нет выдаст ноль

	//как определить есть ли ключ
	val, exists := m["str"]
	if exists {
		fmt.Println(val)
	}

	delete(m, "orange")
}
