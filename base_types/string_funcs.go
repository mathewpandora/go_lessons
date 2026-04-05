package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//преобразование числа в строку
	i := 555
	var str = string(i) //он думает что мы передаем идентификатор символва
	fmt.Println(str)
	s1 := strconv.Itoa(i) //переделывает int в стр по питоновски
	fmt.Println(s1)
	s2 := strconv.FormatInt(int64(i), 2) // строка в двоичном виде
	fmt.Println(s2)
	fmt.Print(reflect.TypeOf(s2))
	number := 10.524986
	fixed := strconv.FormatFloat(number, 'f', 2, 64) //1 - сам флоат 2-читать доку 3-сколько после точки 4-сколько бит в фолте)
	fmt.Println(fixed)

	str_ := fmt.Sprintf("Вот число: %v", fixed)
	fmt.Println(str_)
}
