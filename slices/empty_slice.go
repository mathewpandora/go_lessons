package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s []string //тут пока nil
	fmt.Println(s == nil, reflect.TypeOf(s))

	s2 := []string{"1", "2"}
	fmt.Println(s2 == nil, reflect.TypeOf(s2))

	//слайсы можем срнавить только с нил

	var s3 []string = nil
	fmt.Println(s3)

}

//Верно. Когда слайс равен nil, это означает, что он не указывает на какой-либо массив, при этом его длина и емкость будут равны 0.
//и объявлении переменной без инициализации, она получает значение nil. В Go nil для слайсов означает, что они не указывают на какой-либо массив.
