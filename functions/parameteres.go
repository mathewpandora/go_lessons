package main

import (
	"fmt"
	"math"
)

var username = "Name"

func main() {
	math.Sqrt()
	sayHello1()
	sayHello2("ИМЯ1", "РОЛЬ1")
}

func sayHello1() { //un, role string   так можно пеедавать параметры с одинаковыми типами
	fmt.Printf("Your name is %s", username) //тут не упадет смотрим по другим областям видимости
}

func sayHello2(un string, role string) { //un параметр функции аргумент передается в параметр
	fmt.Printf("Your name is %s and your role is %s", un, role)
}
