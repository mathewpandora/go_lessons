package main

import "fmt"

func main() {
	var p *int                               //объявили переменнуб p в котором должен лежать указатель
	fmt.Println("По умолчаниб в p лежит", p) //nil
	//nil - отсутсвие значения
	p2 := new(int) // new - создаст новый указатель для типа инт
	fmt.Println(p2)
}
