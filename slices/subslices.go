package main

import "fmt"

func main() {
	orig := []int{1, 2, 3, 4, 5}
	subSlice := orig[1:4] //[1:4)
	fmt.Println(subSlice)
	//со строками так де работает
	str := "DFVDFDFGV__укапмувкеп"
	str1 := str[10:] //тут мы взяли байты вроде как могут быть сбои (мы берем не руны а байты а некоторыве симвлы могут занимать несколько байтов )
	fmt.Println(str1)

	//в срезах не копии а оригнапльнвын данные поэтому если поменяем значение в подслайсе в оригиналне тоде изменитс
	orig1 := []int{1, 2, 3, 4, 5}
	orig2 := make([]int, 0, 50)
	fmt.Printf("orig1: %v\n", orig1)
	fmt.Printf("orig2: %v\n", orig2)
	fmt.Println(orig2[2:35]) //ошибки не будет потому большоцй копасити
	//ошибка - fmt.Println(orig1[2:35])
	fmt.Println(orig1[0:1])
}
