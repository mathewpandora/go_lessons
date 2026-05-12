package main

import "fmt"

type UserIdStr int

type ProductIdStr int

type userStr struct {
	ID        UserIdStr
	Firstname string
}

func (u userStr) String() string {
	return fmt.Sprintf("User: %s Id: %d", u.Firstname, u.ID)
}

type ProductStr struct {
	ID    ProductIdStr
	Title string
}

func (p ProductStr) String() string {
	return fmt.Sprintf("User: %s Id: %d", p.Title, p.ID)
}

func main() {

	user1 := userStr{
		ID:        5,
		Firstname: "Ivan",
	}

	product1 := ProductStr{
		ID:    2,
		Title: "Cofee",
	}

	//user1.ID = product1.ID так делать нельзя надо заменить (просто добавили пасеводнимы типам айди и теперь нельзя их присваивать )

	fmt.Println(user1, product1)
}
