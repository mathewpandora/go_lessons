package main

import "fmt"

type Order struct {
	Uuid string
	Cost int
}

func (o Order) Delete() {
	fmt.Println("Order %s (%d rubs) delete")
}

func (o Order) Sale() {
	o.Cost -= 100
}

func main() {
	order1 := Order{
		Uuid: "1111-ffff-rrrr-zzzz",
		Cost: 453234,
	}
	order1.Delete() //данные тут передаются копии и если поменяем ничеего не изменится

	fmt.Println(order1)
}
