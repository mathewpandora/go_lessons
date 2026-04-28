package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	text := "текст, текст, новый тест!"

	text = strings.Map(deletePunct, text) //в функцию отдается каждый элемент строки по очредери

	words := strings.Split(text, " ")

	counter := make(map[string]int)
	for _, word := range words {
		counter[word] += 1
	}

	fmt.Println(counter)
}

func deletePunct(r rune) rune {
	if unicode.IsPunct(r) {
		return -1 //любое отрицательное значение выбросит символ строки
	}
	return unicode.ToLower(r)
}
