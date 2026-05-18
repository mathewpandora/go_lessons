package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

// Оборачивание ошибок
func main() {

	res, err := calculate(6, 0)
	if err != nil {
		log.Fatalf("calculate: %v", err)
	}

	fmt.Println(res)

}

func calculate(a, b int) (int, error) {
	if rand.Intn(100) > 50 {
		return 0, errors.New("ОШибка")
	}
	res, err := divide(a, b)
	if err != nil {
		return 0, fmt.Errorf("divide %w", err)
	}
	return res, nil
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Divide by zero")
	}
	return a / b, nil
}
