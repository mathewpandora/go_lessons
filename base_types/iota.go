package main

import "fmt"

const ( //заполняет константы начиная с 0 на 1
	red = iota
	green
	blue
)

const (
	red1 = iota //обнулся и опять пойдет с 0
	green1
	blue1
)

const ( //тут с 0 по порядку
	a1 = iota
	b1
	c1
	a2 = iota
	b2
	c2
)

const ( //c 3x
	_ = iota + 3 //что такое iota - вообще как сущность ?
	B
	_
	C
)

func main() {
	fmt.Println(red, green, blue)
	fmt.Println(red1, green1, blue1)
	fmt.Println(a1, b1, c1, a2, b2, c2)
	fmt.Println(B, C)


	//iota — это специальный счетчик внутри блока const
	//он сбрасывается заново в каждом новом блоке const
//iota — это специальный счетчик внутри блока const. В начале блока он равен 0 и увеличивается на 1 на каждой следующей строке. Часто его используют, чтобы удобно задавать набор последовательных констант.`
}
