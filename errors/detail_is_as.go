package main

import (
	"errors"
	"fmt"
	"log"
)

// потом узнать почему нельзя записать ошибку в коснтсатну
var My_Error = errors.New("Error")

func main() {
	//fmt.Errorf - отдает указатель на структуру WrapError. or WrapErros
	//error - это интерфейс с методом Error
	//WrapError у нее есть метод UnWrap - который отдает  WrapError.err
	//UnWrap - метод для того чтобы работали  Is  As

	//error Is/As бдует работать с Error.New()

	if err := fn_1(); err != nil {
		if errors.Is(err, My_Error) {
			log.Fatal("1")
		} else {
			log.Fatal("2")
		}
	}
}

func fn_1() error {
	return fmt.Errorf("fn2: %w", fn_2())
}

func fn_2() error {
	return My_Error
}
