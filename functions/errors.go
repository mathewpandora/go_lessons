package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	//в го ошибки отдаются в виде значений
	res, err := devide(5, 2)
	fmt.Println("Резульатат первого вызова", res, err)
	res, err = devide(5, 0)
	fmt.Println("Результат выторого вызова", res, err)
	if err != nil {
		log.Fatalf("divide error %s", err) //Printf b os.Exit(1)
	}

}

func SomeFun1() error {
	fmt.Errorf("make %s", "somefunc") //тоже отдает ошибку
	return errors.New("Some error")
}

func devide(n1, n2 int) (int, error) {
	if n2 == 0 {
		return 0, errors.New("devide by zero")
	} else {
		return n1 / n2, nil
	}
}
