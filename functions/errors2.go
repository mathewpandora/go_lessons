package main

import (
	"fmt"
	"log"
)

func main() {
	a, b := 10.0, 2.0
	res, err := calculate(a, b)
	if err != nil {
		log.Fatalf("unable to calculate $s", err)
	}
	fmt.Println("result:", res)
}

func calculate(a, b float64) (float64, error) {
	if err := logicX(); err != nil {
		return 0, err
	}
	return divide(a, b)
}

func logicX() error {
	if err := logicY(); err != nil {
		fmt.Errorf("logicY %w", err) // правильный метод ошибку описывать чтобвы вдитель в какой конкретноц функции она произошла
	} //вотт так стоим тоборачивать только в случяаве когда орбабатываем ошибку
	return nil
}

func logicY() error {
	return nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("b == 0")
	}
	return a / b, nil
}

//log.Fatalf - падает с ошбикой
//fmt.Errorf - использовать в функцичх
//errors.New - ошибка сама

// Узнать про них более подробно в чем разнца и в чем отличия
//и еще раз в чем рахзница между паникой и ошибкой
