package main

import "fmt"

func main() {
	var day int
	fmt.Print("День от 1 до 7:")
	fmt.Scan(&day)

	//если в кейч положить переменную - есть варик упасть поттму что го не отслеживает что лежит в переменныъ -> почериквать ошибку не будет
	//когда кейс встретиля - дальше не продолжаем проверять а сразу уходим из блока
	switch day {
	case 1:
		fmt.Println("Понедельник")
		fallthrough //зайди в следующий кейс
	case 2:
		fmt.Println("Вторник")
	case 3:
		fmt.Println("среда")
	case 4:
		fmt.Println("Четверг")
	case 5:
		fmt.Println("Пятница")
	case 6:
		fmt.Println("Суббота")
	case 7:
		fmt.Println("Воскрсенсене")
	case 8, 9, 10:
		fmt.Println("Воскрсенсене")
	default:
		fmt.Println("Нет такого дня")
	}

	switch {
	case day == 1:
		fmt.Println("Понедельник")
	case day == 2:
		fmt.Println("Вторник")
	}

	//для орпедления типа данных:
	var x any = "sad"

	switch v := x.(type) { //.(type) — специальное ключевое слово, которое работает только внутри switch v — переменная, которая получает значение x, но уже приведённое к конкретному типу
	case int:
		fmt.Println("инт", v)
	case string:
		fmt.Println("стринг", v)
	case bool:
		fmt.Println("бул", v)
	default:
		fmt.Println("хз", v)
	}
}
