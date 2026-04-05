package main

import "fmt"

func main() {
	var (
		name      string
		lastnamne string
		age       int
	)
	fmt.Print("Введите ваше имя и фамилию: ")
	fmt.Scanln(&name, &lastnamne)
	fmt.Print("Введите ваш возраст: ")
	fmt.Scanln(&age) // на ошибку пока не проверяю
	fmt.Printf("Приятно познакомиться, %s. Я 5 лет назад познакомился с человеком, у которого тоже фамилия %s, вам тогда было %d. Как молоды мы были!", name, lastnamne, age-5)
}
