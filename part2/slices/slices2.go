package main

import "fmt"

func main() {
	//	append го проверяет хвтит ли места чтобы вставить в массив
	// 	if len(newElems) > cap(s) - len(s) - требуется релокация
	//	релокация - массив x2 - пересоздание масива она дорогая по ремени и по памяти

	v := make([]int, 2, 3)

	fmt.Println(cap(v), len(v))

	v = append(v, 10, 20)

	fmt.Println(cap(v), len(v))

	//пока слайсы маленькие их рост при релокации будет x2
	//если слайс большой при релокации нового масива будет не так сильно уведичиваться массив (бдует где то на 25% увеличиваться)
	//https://github.com/golang/go/tree/master/src/runtime
	//https://github.com/golang/go/blob/master/src/runtime/slice.go
	//nextSliceCap - функция которая определет какой будет следуюий капасити у слайса (из доки)
	//newcap = (newcap * newcap + threshold 3*(256) ) / 4


}
