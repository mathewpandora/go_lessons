package main

import "fmt"

func main() {
	var val interface{} = "что угошдно?" //можем положить что угодно пустой интерфейс
	fmt.Println(val)
	val = 23
	fmt.Println(val)
	var val2 any = "что угошдно?" //то же самое
	fmt.Println(val2)
}
