package main

import "fmt"

type Order struct {
	Uuid string
	Cost int
}

func (o Order) Delete() {
	fmt.Println("Order %s (%d rubs) delete")
}

func (o Order) Sale1() {
	o.Cost -= 100
}

func (o *Order) Sale2() {
	o.Cost -= 100
}

func (Order) SayHello() string {
	return "hello" //тут если неиррспольузем атрибуты структуры
}
func main() {
	order1 := Order{
		Uuid: "1111-ffff-rrrr-zzzz",
		Cost: 453234,
	}
	//тут данные не мутруются
	order1.Sale1() //данные тут передаются копии и если поменяем ничеего не изменится

	fmt.Println(order1)
	//тут данные мутриуются
	order1.Sale2() //данные тут передаются по указателю -> изменятся данные

	fmt.Println(order1)
}
