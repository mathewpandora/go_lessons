package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Weight float64
	Height float64
}

type Cirle struct {
	Radius float64
}

func (c Cirle) FindSquare() float64 {
	return c.Radius * math.Pi
}

func (c Rectangle) FindSquare() float64 {
	return c.Height * c.Weight
}

type FigureSquareFinder interface {
	FindSquare() float64
}

func main() {

	var a any //any - это пустой интерфейс
	fmt.Println(a)

	s := []any{
		Cirle{Radius: 54},
		Rectangle{Weight: 80, Height: 70},
	}

	PrintFigure(s)

	//но можно сделать слайс интерфейсов (структур с общим методом)

	//интерфейс может быть нил ! Это всегда надо учитывать !
}

func AnyPrint(a any) {
	fmt.Println(a)
}

func PrintFigure(s []any) {
	for _, v := range s {
		if s, ok := v.(FigureSquareFinder); ok { //Это проверка типа (type assertion) с проверкой успеха, которая в Go называется "comma ok" idiom.
			fmt.Printf("Figure Data : %v\n", s.FindSquare())
		} else if s, ok := v.(string); ok {
			fmt.Println("Это не строка", s)
		}
	}

	//var b1 interface {
	//}

	//b1 = 2 //


	//var b2 any

	//b2 = 3

	//fmt.Println(b2 + 4)

	//fmt.Printf("%t", b1 + 2)
}

//Да, интерфейс считается nil, если он не содержит ни значения, ни типа. Это означает, что он не инициализирован и не указывает на какой-либо объект.
//Интерфейс в Go — это структура из двух полей:

//Указатель на тип (динамический тип)
//Указатель на значение (динамическое значение)

//я не пончял что значит что в интерфейцсе лежит указатель на тип и указатель нап значение
