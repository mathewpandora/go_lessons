package main

import (
	"fmt"
	"strings"
)

func main() {
	//есть функция strings.Contains - проверяет вхождение подстроки в строку
	str := "	fvcwc 	"
	fmt.Println(strings.TrimSpace(str))
	fmt.Println(str)
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
	str2 := "word1,word2,word3"
	fmt.Println(strings.Split(str2, ","))
	fmt.Println(strings.Contains(str2, "w"))  //содержит
	fmt.Println(strings.HasPrefix(str2, "w")) //начинается ли
	fmt.Println(strings.HasSuffix(str2, "w")) //оканчивается ли
}
