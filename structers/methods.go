package main

import "fmt"

type Car struct {
	Model string
	Speed int
}

func (u Car) Drive() { //u - то к какой структуре относится метод (метод отонсится к типу )
	fmt.Printf("Автомобиль %s поехал со скростью %d", u.Model, u.Speed)
}

func main() {

	car := Car{
		Model: "BMW",
		Speed: 93,
	}

	car.Drive()

	car.Model = "Nissan"

	car.Drive()
}
