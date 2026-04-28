package main

import "fmt"

type User4 struct {
	Name string
}

func main() {

	user := User4{
		Name: "Matvey",
	}

	modifyValue1(user)

	fmt.Println(user)
}

func modifyValue1(u User4) {
	u.Name = "fgbvrt" //тут передалась копия помиенялось только тут а оригналтне поменялся
}

func modifyValue2(u *User4) { //принимаем указатель
	*u.Name = "fgbvrt"
}
