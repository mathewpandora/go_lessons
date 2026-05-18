package main

import (
	"errors"
	"fmt"
)

func main() {

}

func fn1() error {
	//error - это интерфейс в котором есть метод Error и он отдает строку
	return errors.New("ошиька")
}

func fn2() error {
	//error - это интерфейс в котором есть метод Error и он отдает строку
	return fmt.Errorf("Ошибка номер %d", 12) //Формирует строку
}
