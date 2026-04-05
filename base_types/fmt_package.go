package main

import "fmt"

func main() {
	fmt.Print("Ghdftn")
	fmt.Print("Ghdftn")
	fmt.Printf("Приает %s %s", "еуыен", "missiun\n") //printf ! %-куда вставляем и потом какой тип
	digit := 123
	fmt.Printf("Число: %d", digit)

	insert1 := "Name"
	insert2 := 12
	str := fmt.Sprintf("Меня зовут %s мне %d лет", insert1, insert2)
	fmt.Println(str)
	//СПЕЦИФИКАТОРЫ f-функций (начинаются с процентов)
	// %s  - string
	// %d - int
	// %f - float
	// %t - bool
	// %% - вставит процент
	// %b - выведет число в двоичной
	//%.2f - флоат два знака после запятой
	//%v - вставляет че по кайфу
	//Восьмеричная	%o
	//Шестнадцатеричная	%x
	//Шестнадцатеричная (верхний регистр)	%X
	//Шестнадцатеричная (с префиксом)	%#x

}
