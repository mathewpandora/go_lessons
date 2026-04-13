package main

import "fmt"

func main() {
	//паника - такая ошибка когда голанг не может работать
	// panic("ПАНИКА!")
	defer handlePanic()
	riskyFunction()
	fmt.Println("Эта строка не будет достигнута, если произойдет паника")

}

func riskyFunction() {
	panic("что то пошлр не так")
}

func handlePanic() {
	if r := recover(); r != nil { //if переменная := выражение; условие {
		fmt.Println("Обработана паникa", r) // с помощью recover можно поймать панику и работает тольуо внутри defer
	} //recover должен вызываться внутри отложенной функции (defer), чтобы корректно перехватить панику.
}
