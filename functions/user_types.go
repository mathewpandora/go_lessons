package main

import "fmt"

// Можем создать новый тип на основании другого типа
type Direction int //отдельный тип точно такой же как int
const (
	_ Direction = iota //заготовленные константы для типа дирекшен
	North
	West
	South
	East
)

func main() {

	action(South)
	action(5) //так будет работать
	dir := 5
	// action(dir) - тут будет ошибка (потому что передавем тип инт)
	fmt.Println(dir)
	action(Direction(dir))
}

func action(d Direction) { //хочется передавать число от 1 до 4 которое будет отдавать направление света (смервер юг вотсок запаД)
	fmt.Println("Дейсвтие в направлении:", d)
}
