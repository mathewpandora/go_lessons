package main

import "fmt"

func main() {

	str := "vbddfgbvmlkvmdfvmndfklvjnm"

	for i, r := range str { //i - индекс первого байта текущей , руны r - руна
		fmt.Println(i, r)
		fmt.Println(string(r))
	}

	for i := 0; i < len(str); i++ {
		fmt.Println(i, str[i]) //выведет все байты строки каждая итерация - один байт
	}
	//НАДО ПОПОДРОБНЕЕ РАОБЗРАТЬСЯ С БАЙТАМИ СТРОКАМИ ЧТО КУДА И ЧЕГО - ТЯЖЕЛО ПОСЛЕ ПИТОНА ДЛЯ ПОНИМАНИЯ

	for _, b := range []byte(str) { //i - индекс байта b - сам байт
		fmt.Println(b)
	}
}
