package main

import "fmt"

type Name struct {
	First string
	Last  string
}

type User2 struct {
	Name      Name
	BirthYear int //композиция
}

//Можнор сделдть анонимную структур у

type User3 struct {
	Name struct {
		First string
		Last  string
	}
}

func main() {

	name := Name{
		First: "Name",
		Last:  "Surname",
	}
	new_user := User2{
		Name:      name,
		BirthYear: 1991,
	}

	fmt.Println(new_user.Name.First)
}
