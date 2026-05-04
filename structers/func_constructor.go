package main

import (
	"fmt"
)

type Person struct {
	Firstname string
	Age       int
	Lastname  string
}

func main() {
	p1 := Person{
		Firstname: "Name",
		Age:       34,
	}

	p2 := Person{
		Firstname: "Name2",
		Age:       342,
	}

	fmt.Println(p1, p2)

	//а ВДруг потом в структуру добавили фамлмию? тогда везде будет зиро валуе
}

// из конерструктор принято отдавать указатель
func NewPerson(firstname string, age int, lastname string) (*Person, error) { //конструкртор создает новаый тип стурктуры
	return &Person{ //Почему возвращаем именно указатель
		Firstname: firstname,
		Age:       age, //отдает указатель
		Lastname:  lastname,
	}, nil
}

//Еще раз, коротко. Из функции-конструктора возвращаем указатель на структуру если:

//Структура огромная (прям очень-очень большая, это очень редко бывает).
//В структуре могут меняться данные.
//Если структура нужна для того чтобы передать данные из одной функции в другую (такие структуры называют DTO, с англ. - Data Transfer Object), в таком случае возвращаем структуру из функции-конструктора (не указатель).
