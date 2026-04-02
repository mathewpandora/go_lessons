package main

import "fmt"

// статическая типизация - объявляем какой тип будет лежать в перемнной
// если переменная не используется потом нигде будет ошибка ./main.go:5:6: declared and not used: ag
func main() {
	var age1 int //создать  переменную (по умолчанию у типа инт - 0)
	var age2 int = 30
	var age3 = 30 //можно не указывать тип он сам поцмет
	age4 := 12    //создать и сразу присвоить
	//В коде нельзя дважды создавать переменную с одним и тем же типом - бдует ошибка
	var x, y int = 1, 2
	x1, y2 := 3, 4

	var (
		a = 5
		b = 2
		c int
	)

	//перезаписать переменную:

	num1 := 24
	fmt.Println(num1)
	num1 = 12
	fmt.Println(num1)
	// констатны:

	const Pi = 3.14

	const (
		//тут поставить просто имя константы - не можем будет ошибка
		a1 = 1
		b1 = 2
		c1 //Для константы по умолчанию кладется значение предыдущей константы
	)

	/*
		тест
	*/
	fmt.Println(age1)
	fmt.Println(age2)
	fmt.Println(age3)
	fmt.Println(age4)
	fmt.Println(x, y)
	fmt.Println(x1, y2)
	fmt.Println(a, b, c)
	fmt.Println(Pi)
	fmt.Println(a1, b1, c1)
}
