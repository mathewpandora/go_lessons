package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	var firstname string
	fmt.Scan(&firstname) //передаем адрес чтобы было понятно куда записывать то что ввели (переменная всего лишь ссылается на адрес )
	//тут спросить про память в рамках функции когда уйдем из неен что будет потом
	fmt.Printf("Ваше имя %s\n", firstname)
	//можно вводитль через пробел

	var (
		a string
		b string
	)
	fmt.Scan(&a, &b)
	fmt.Printf("Ввели %s and %s", a, b)
	//fmt.Scanln() - для чтения строки:
	//как прочитать всю строку:
	scanner := bufio.NewScanner(os.Stdin) //аргумент - откуда он считывает текст os.Stdin - терминальчик  standard input
	fmt.Println("Тип сканера:", reflect.TypeOf(scanner))
	fmt.Print("Введи строку:")
	scanner.Scan()
	fmt.Println("Вы ввели:", scanner.Text())
}
