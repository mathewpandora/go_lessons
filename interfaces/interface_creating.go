package main

import "fmt"

type Cat1 struct {
	name string
}

func (c Cat1) Growl() string {
	return "sdefvdfv1"
}

type Cat2 struct {
	name string
}

func (c Cat2) Growl() string {
	return "sdefvdfv2"
}

func main() {
	//интерфейс - это тип который определяент набор методов
	c1 := Cat1{
		name: "1",
	}

	c2 := Cat2{
		name: "2",
	}

}

type Growler interface {
	Growl() string //структуры
}

// утиная типизация - если что крякает как утра - это утка
func makeSound(g Growler) { //Тип у которого нет метода будет ошибка (это полимофризм буквальео единый интерфейс )
	fmt.Println(g.Growl())
}
