package main

import "fmt"

type User struct { //структура (как класс) это типа наш тип
	Firstname string
	Lastname  string
	BirthYear int
	IsAdmin   bool
}

func main() {
	//иногда надо хранить разные данные для одной сущности
	//Структуры создаются с помощьб type

	new_user := User{"Dima", "Familia", 1991, true}
	fmt.Println(new_user)
	new_user2 := User{Firstname: "Dima", Lastname: "Familia", BirthYear: 1991, IsAdmin: true}

	fmt.Println(new_user2.Firstname)

	new_user3 := User{} //тут зиро вэлуе лежит

	var new_user4 User //тут зиро валуе все по типапм

	fmt.Println(new_user3, new_user4)

	new_user5 := &User{} //тут указатель на структру ё

	fmt.Println(new_user3, new_user4, new_user5)
}
