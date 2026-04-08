package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	res := find_sum1(2, 3)
	fmt.Println(res)
	str := "gbfggbgfb вамвамвамавм"
	_, runes := getFullLenght1(str) // _ не требуется для использования
	fmt.Println(runes)
}

func find_sum1(n1, n2 int) int {
	return n1 + n2
}

func getFullLenght1(str string) (bytes int, runes int) {
	//Именованные возвращаемые значения всегда инициализируются значениями по умолчанию при создании, поэтому ошибка не произойдет, если они не будут явно инициализированы внутри функции.
	bytes = len(str)
	runes = utf8.RuneCountInString(str)
	return // автоматом вернутся переменные  bytes int, runes int
}
