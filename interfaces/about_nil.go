package main

import "fmt"

type Cruder interface {
	Create() bool
	Delete() bool
	Read() bool
	Insert() bool
}

func main() {
	//nil

	var my_nil *int //по умолчанию нил
	//Важно что под слайс мордно тодже передать нил и тогда если пользвоатсья им как слайссом будет ошибка
	//но с меп так не работает даже если он нил - мы можем читать ключи которых нет и будет ноль (но только если там должен лежать меп)
	//
	if my_nil == nil {
		fmt.Println("slice is nill")
	}

	res := DoCrud(nil)
	fmt.Println(res)
}
func DoCrud(c Cruder) bool {
	if c == nil {
		fmt.Println("NIL!") //короче если принимает какой-то интерфейс то принимаем нил!
		return false
	} else {
		fmt.Println("NIL!")
		return false
	}
}
