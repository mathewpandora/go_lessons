package main

import "fmt"

type Pers struct {
	Name string
	Age  int
}

func (p Pers) SayHello() {
	fmt.Printf("%s said Hello", p.Name)
}

func (p Pers) Struct() {
	fmt.Printf("Person ")
}

func (p *Pers) StructPointer() {
	fmt.Printf("Person")
}

func main() {

	var p *Pers //переменая нил в ней должен лежать указатель на перс

	//fmt.Println(p.SayHello())
	//fmt.Println(p.Struct())- ошибка потому что эта запись эквыивалента *p(получили значение = nil).Struct() *p (получает по )
	p.StructPointer() //так отработает потому что метод ожидает у4казатель

}
